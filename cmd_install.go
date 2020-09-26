package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
	"github.com/mholt/archiver/v3"
)

type Release struct {
	URL            string
	Owner          string
	Repo           string
	Version        string
	AssetFileName  string
	InstalledFiles InstalledFiles
}
type Releases []Release

type InstalledFile struct {
	Src, Dest string
}
type InstalledFiles []InstalledFile

func (a *App) CmdInstall(url string) error {
	rel, err := parseURL(url)
	if err != nil {
		return err
	}

	dir := a.Config.ReleasesDir()
	releasesDir := filepath.Join(dir, rel.Owner, rel.Repo, rel.Version)
	if err := os.MkdirAll(releasesDir, os.ModePerm); err != nil {
		return err
	}

	assetFile, err := downloadFile(rel.URL, releasesDir, rel.AssetFileName)
	if err != nil {
		return err
	}

	assetDir := filepath.Join(releasesDir, "assets")
	if err := os.MkdirAll(assetDir, os.ModePerm); err != nil {
		return err
	}
	if err := archiver.Unarchive(assetFile, assetDir); err != nil {
		return err
	}

	binDir := a.Config.BinDir()
	installedFiles, err := installFiles(assetDir, binDir)
	if err != nil {
		return err
	}
	rel.InstalledFiles = installedFiles

	var rels Releases
	releasesFile := filepath.Join(a.Config.RelmaRoot, "releases.json")
	_, err = os.Stat(releasesFile)
	if os.IsExist(err) {
		b, err := ioutil.ReadFile(releasesFile)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(b, &rels); err != nil {
			return err
		}
	}
	rels = append(rels, *rel)
	b, err := json.MarshalIndent(&rels, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(releasesFile, b, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func parseURL(s string) (*Release, error) {
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

	p := &Release{
		URL:           s,
		Owner:         owner,
		Repo:          repo,
		Version:       version,
		AssetFileName: file,
	}

	return p, nil
}

func downloadFile(url, destDir, destFile string) (string, error) {
	resp, err := http.Get(url)
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

func installFiles(srcDir, destDir string) (InstalledFiles, error) {
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

	if binDir == "" {
		return ifs, nil
	}

	ifs2, _, err := linkExecutableFilesToDest(binDir, destDir)
	if err != nil {
		return nil, err
	}

	ifs = append(ifs, ifs2...)

	return ifs, nil
}

func linkExecutableFilesToDest(srcDir, destDir string) (InstalledFiles, string, error) {
	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return nil, "", err
	}

	var ifs InstalledFiles
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

func linkExecutableFileToDest(f os.FileInfo, src, dest string) (*InstalledFile, error) {
	isExec, err := isExecutableFile(f, src)
	if err != nil {
		return nil, err
	}
	if !isExec {
		return nil, nil
	}

	if err := os.Symlink(src, dest); err != nil {
		return nil, err
	}
	ff := InstalledFile{
		Src:  src,
		Dest: dest,
	}
	return &ff, nil
}

func isExecutableFile(f os.FileInfo, path string) (bool, error) {
	mode := f.Mode()
	if !mode.IsRegular() {
		return false, nil
	}

	if mode&0111 != 0 {
		return true, nil
	}

	typ, err := filetype.MatchFile(path)
	if err != nil {
		return false, err
	}
	switch typ.Extension {
	case "elf", "exe":
		return true, nil
	}

	return false, nil
}
