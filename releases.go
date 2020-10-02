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

func (r *Release) FormatSimpleInformation() string {
	return fmt.Sprintf("%s/%s %s", r.Owner, r.Repo, r.Version)
}
