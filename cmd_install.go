package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
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

	assetFile, err := downloadFile(pkg.URL, pkgDir)
	if err != nil {
		return err
	}

	assetDir := filepath.Join(pkgDir, "assets")
	if err := unarchiveFile(assetFile, assetDir); err != nil {
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

func parseURL(url string) (Package, error) {
	return Package{}, nil
}

func downloadFile(url, destDir string) (string, error) {
	return "", nil
}

func unarchiveFile(path, destDir string) error {
	return nil
}

func installFiles(dir string) (InstalledFiles, error) {
	return InstalledFiles{}, nil
}
