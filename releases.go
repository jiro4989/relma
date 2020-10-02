package main

import (
	"fmt"
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

func (files InstalledFiles) FixPath(srcDir, destDir string) {
	for i := 0; i < len(files); i++ {
		file := &files[i]
		file.Src = file.Src[len(srcDir)+1:]
		file.Dest = file.Dest[len(destDir)+1:]
	}
}
