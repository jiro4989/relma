package main

import (
	"errors"
	"fmt"
	"strings"
)

type Release struct {
	URL            string
	Owner          string
	Repo           string
	Version        string
	LatestVersion  string
	AssetFileName  string
	InstalledFiles InstalledFiles
}

type Releases []Release

type InstalledFile struct {
	Src, Dest string
}

type InstalledFiles []InstalledFile

func (r *Release) FormatSimpleInformation() string {
	return fmt.Sprintf("%s/%s %s", r.Owner, r.Repo, r.Version)
}

func (r *Release) EqualRepo(ownerRepo string) (bool, error) {
	oRepo := strings.Split(ownerRepo, "/")
	if len(oRepo) < 2 {
		msg := fmt.Sprintf("%s is illegal format", ownerRepo)
		return false, errors.New(msg)
	}

	ok := strings.ToLower(oRepo[0]) == strings.ToLower(r.Owner) && strings.ToLower(oRepo[1]) == strings.ToLower(r.Repo)
	return ok, nil
}

func (files InstalledFiles) FixPath(srcDir, destDir string) {
	for i := 0; i < len(files); i++ {
		file := &files[i]
		file.Src = file.Src[len(srcDir)+1:]
		file.Dest = file.Dest[len(destDir)+1:]
	}
}
