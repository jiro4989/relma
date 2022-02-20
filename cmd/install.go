package cmd

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/jiro4989/relma/filetype"
	"github.com/jiro4989/relma/releases"
	"github.com/mholt/archiver/v3"
	"github.com/spf13/cobra"
)

func init() {
	commandInstall.Flags().StringVarP(&commandLineInstallParam.File, "file", "f", "", "install with releases.json")

	rootCmd.AddCommand(commandInstall)
}

type CommandLineInstallParam struct {
	GitHubReleaseURL string `docopt:"<github_release_url>"`
	File             string `docopt:"-f,--file"`
}

var commandLineInstallParam CmdInstallParam

var commandInstall = &cobra.Command{
	Use:   "install",
	Short: "install GitHub Releases",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := NewApp()
		if err != nil {
			return err
		}
		commandLineInstallParam.URL = args[0]
		return a.CmdInstall(&commandLineInstallParam)
	},
}

type CmdInstallParam struct {
	URL  string
	File string
}

// CmdInstall installs commands from an url of GitHub Releases.
//
// Decompress it if Releases file was archive file, and install executables to
// relma directory.Install executables from urls of `releases.json` if `p.File`
// (releases.json) is not empty.
func (a *App) CmdInstall(p *CmdInstallParam) error {
	if p.File != "" {
		rels, err := ReadReleasesFile(p.File)
		if err != nil {
			return err
		}
		var errCount int
		for _, rel := range rels {
			p := CmdInstallParam{URL: rel.URL}
			err := a.CmdInstall(&p)
			if err != nil {
				Error(err)
				errCount++
			}
			Sleep()
		}
		if 0 < errCount {
			return errors.New("install failed")
		}
		return nil
	}

	url := p.URL
	rel, err := parseURL(url)
	if err != nil {
		return err
	}

	dir := a.Config.ReleasesDir()
	releasesDir := filepath.Join(dir, rel.Owner, rel.Repo, rel.Version)

	// skip if releases directory has already existed
	_, err = os.Stat(releasesDir)
	if !os.IsNotExist(err) {
		msg := fmt.Sprintf("skip: %s has already installed", rel.FormatVersion())
		fmt.Println(msg)
		return nil
	}

	if err := os.MkdirAll(releasesDir, os.ModePerm); err != nil {
		return err
	}

	assetFile, err := a.downloadFile(rel.URL, releasesDir, rel.AssetFileName)
	if err != nil {
		return err
	}

	assetDir := filepath.Join(releasesDir, "assets")
	if err := os.MkdirAll(assetDir, os.ModePerm); err != nil {
		return err
	}

	if ok, err := filetype.IsArchiveFile(assetFile); err != nil {
		return err
	} else if !ok {
		// Download and create symlink if assetFile is not archive file and is
		// executable file. Not unarchive.
		oldAssetFile, assetFile := assetFile, filepath.Join(assetDir, filepath.Base(assetFile))
		if err := os.Rename(oldAssetFile, assetFile); err != nil {
			return err
		}

		fi, err := os.Stat(assetFile)
		if err != nil {
			return err
		}

		if ok, err := filetype.IsExecutableFile(fi, assetFile); err != nil {
			return err
		} else if !ok {
			err = errors.New("github_releases_url file must be executable file or archive file")
			return err
		}

		name := filepath.Base(assetFile)
		binDir := a.Config.BinDir()
		destFile := filepath.Join(binDir, name)
		ff, err := linkExecutableFileToDest(fi, assetFile, destFile)
		if ff == nil && err == nil {
			// 到達しないはず
			err = errors.New("unknown error")
			return err
		}
		if err != nil {
			return err
		}
		rel.InstalledFiles = releases.InstalledFiles{*ff}
		rel.InstalledFiles.FixPath(filepath.Dir(assetFile), binDir)

		rels, err := a.Config.ReadReleasesFile()
		if err != nil {
			return err
		}
		if existed, index := existsRepo(rels, *rel); existed {
			rels[index] = *rel
		} else {
			rels = append(rels, *rel)
		}
		err = a.SaveReleases(rels)
		if err != nil {
			return err
		}

		fmt.Println(rel.FormatVersion())
		return nil
	}

	// Download and archive file and create symlink if assetFile is archive file.

	if err := archiver.Unarchive(assetFile, assetDir); err != nil {
		return err
	}

	installedFiles, err := installFiles(assetDir, a.Config.BinDir())
	if err != nil {
		return err
	}
	rel.InstalledFiles = installedFiles

	rels, err := a.Config.ReadReleasesFile()
	if err != nil {
		return err
	}
	if existed, index := existsRepo(rels, *rel); existed {
		rels[index] = *rel
	} else {
		rels = append(rels, *rel)
	}
	err = a.SaveReleases(rels)
	if err != nil {
		return err
	}

	fmt.Println(rel.FormatVersion())

	return nil
}

func parseURL(s string) (*releases.Release, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	if strings.ToLower(u.Host) != "github.com" {
		return nil, errors.New("only use 'github.com' domain")
	}

	paths := strings.Split(u.Path, "/")
	if len(paths) < 7 {
		return nil, errors.New("illegal install URL")
	}
	owner := paths[1]
	repo := paths[2]
	version := paths[5]
	file := paths[6]

	if owner == "" || repo == "" || version == "" || file == "" {
		return nil, errors.New("illegal install URL")
	}

	p := &releases.Release{
		URL:           s,
		Owner:         owner,
		Repo:          repo,
		Version:       version,
		LatestVersion: version,
		AssetFileName: file,
	}

	return p, nil
}

func (a *App) downloadFile(url, destDir, destFile string) (string, error) {
	resp, err := a.Downloader.Download(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	destPath := filepath.Join(destDir, destFile)
	file, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return destPath, nil
}

func installFiles(srcDir, destDir string) (releases.InstalledFiles, error) {
	df, err := os.Stat(destDir)
	if os.IsNotExist(err) {
		return nil, err
	}

	if !df.IsDir() {
		msg := fmt.Sprintf("'%s' must be directory", destDir)
		return nil, errors.New(msg)
	}

	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return nil, err
	}

	if len(files) == 1 && files[0].IsDir() {
		srcDir = filepath.Join(srcDir, files[0].Name())
	}

	ifs, binDir, err := linkExecutableFilesToDest(srcDir, destDir)
	if err != nil {
		return nil, err
	}
	if ifs != nil {
		ifs.FixPath(srcDir, destDir)
	}

	if binDir == "" {
		return ifs, nil
	}

	ifs2, _, err := linkExecutableFilesToDest(binDir, destDir)
	if err != nil {
		return nil, err
	}
	if ifs2 != nil {
		ifs2.FixPath(srcDir, destDir)
	}

	ifs = append(ifs, ifs2...)

	return ifs, nil
}

func linkExecutableFilesToDest(srcDir, destDir string) (releases.InstalledFiles, string, error) {
	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return nil, "", err
	}

	var ifs releases.InstalledFiles
	var binDir string
	for _, f := range files {
		name := f.Name()
		srcFullPath := filepath.Join(srcDir, name)
		destFullPath := filepath.Join(destDir, name)

		if name == "bin" {
			binDir = srcFullPath
		}

		ff, err := linkExecutableFileToDest(f, srcFullPath, destFullPath)
		if ff == nil && err == nil {
			continue
		}
		if err != nil {
			return nil, "", err
		}

		ifs = append(ifs, *ff)
	}
	return ifs, binDir, nil
}

func linkExecutableFileToDest(f os.FileInfo, src, dest string) (*releases.InstalledFile, error) {
	isExec, err := filetype.IsExecutableFile(f, src)
	if err != nil {
		return nil, err
	}
	if !isExec {
		return nil, nil
	}

	err = os.Chmod(src, f.Mode()|0111)
	if err != nil {
		return nil, err
	}

	if _, err := os.Lstat(dest); err == nil {
		os.Remove(dest)
	}

	if err := os.Symlink(src, dest); err != nil {
		return nil, err
	}
	ff := releases.InstalledFile{
		Src:  src,
		Dest: dest,
	}
	return &ff, nil
}

func existsRepo(rels releases.Releases, rel releases.Release) (bool, int) {
	for i, r := range rels {
		if !r.EqualRelease(&rel) {
			continue
		}
		return true, i
	}
	return false, -1
}
