package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/mholt/archiver/v3"
)

type Package struct {
	URL           string
	Owner         string
	Repo          string
	Version       string
	AssetFileName string
}

type InstalledFiles []InstalledFile
type InstalledFile struct {
	Src, Dest string
}

type PackageInfo struct {
	URL            string
	InstalledFiles InstalledFiles
}

func (a *App) CmdInstall(url string) error {
	pkg, err := parseURL(url)
	if err != nil {
		return err
	}

	dir := a.Config.PackageDir()
	pkgDir := filepath.Join(dir, pkg.Owner, pkg.Repo, pkg.Version)
	if err := os.MkdirAll(pkgDir, os.ModePerm); err != nil {
		return err
	}

	assetFile, err := downloadFile(pkg.URL, pkgDir, pkg.AssetFileName)
	if err != nil {
		return err
	}

	assetDir := filepath.Join(pkgDir, "assets")
	if err := archiver.Unarchive(assetFile, assetDir); err != nil {
		return err
	}

	installedFiles, err := installFiles(assetDir)
	if err != nil {
		return err
	}

	p := PackageInfo{
		URL:            pkg.URL,
		InstalledFiles: installedFiles,
	}
	b, err := json.Marshal(&p)
	if err != nil {
		return err
	}

	pkgFile := filepath.Join(pkgDir, "pkginfo.json")
	if err := ioutil.WriteFile(pkgFile, b, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func parseURL(s string) (*Package, error) {
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

	p := &Package{
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

func installFiles(dir string) (InstalledFiles, error) {
	return InstalledFiles{}, nil
}
