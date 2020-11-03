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

func RemoveRelease(rels Releases, rel *Release) Releases {
	for i, v := range rels {
		if !rel.EqualRelease(&v) {
			continue
		}
		return unset(rels, i)
	}
	return rels
}

func unset(s Releases, i int) Releases {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

func (r *Release) FormatSimpleInformation() string {
	return fmt.Sprintf("%s/%s %s", r.Owner, r.Repo, r.Version)
}

func (r *Release) FormatVersion() string {
	return fmt.Sprintf("%s/%s:%s", r.Owner, r.Repo, r.Version)
}

func (r *Release) EqualRepo(ownerRepo string) (bool, error) {
	oRepo := strings.Split(ownerRepo, "/")
	if len(oRepo) < 2 {
		msg := fmt.Sprintf("%s is illegal format", ownerRepo)
		return false, errors.New(msg)
	}

	ok := strings.EqualFold(oRepo[0], r.Owner) && strings.EqualFold(oRepo[1], r.Repo)
	return ok, nil
}

func (r *Release) EqualRelease(r2 *Release) bool {
	ok := strings.EqualFold(r.Owner, r2.Owner) && strings.EqualFold(r.Repo, r2.Repo)
	return ok
}

func (files InstalledFiles) FixPath(srcDir, destDir string) {
	for i := 0; i < len(files); i++ {
		file := &files[i]
		file.Src = file.Src[len(srcDir)+1:]
		file.Dest = file.Dest[len(destDir)+1:]
	}
}
