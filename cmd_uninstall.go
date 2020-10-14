package main

import (
	"errors"
	"os"
	"path/filepath"
)

type CmdUninstallParam struct {
	OwnerRepo string
}

func (a *App) CmdUninstall(p *CmdUninstallParam) error {
	rels, err := a.Config.ReadReleasesFile()
	if err != nil {
		return err
	}

	rel, err := uninstallableRelease(rels, p.OwnerRepo)
	if err != nil {
		return err
	}

	_, err = a.uninstallRelease(rel, p)
	if err != nil {
		return err
	}

	rels = RemoveRelease(rels, rel)
	err = a.SaveReleases(rels)
	if err != nil {
		return err
	}

	MessageOK("uninstall", rel.FormatVersion())

	return nil
}

func uninstallableRelease(rels Releases, ownerRepo string) (*Release, error) {
	var rel *Release
	for _, r := range rels {
		if ok, err := r.EqualRepo(ownerRepo); err != nil {
			return nil, err
		} else if !ok {
			continue
		}
		rel = &r
		break
	}
	if rel == nil {
		return nil, errors.New(ownerRepo + " was not installed")
	}
	return rel, nil
}

func (a *App) uninstallRelease(rel *Release, p *CmdUninstallParam) ([]string, error) {
	var removedFiles []string
	releaseDir := filepath.Join(a.Config.ReleasesDir(), rel.Owner, rel.Repo)
	err := os.RemoveAll(releaseDir)
	if err != nil {
		return nil, err
	}
	removedFiles = append(removedFiles, releaseDir)

	var errs []error
	for _, f := range rel.InstalledFiles {
		d := filepath.Join(a.Config.BinDir(), f.Dest)
		err := os.Remove(d)
		if err != nil {
			MessageNG("uninstall", "failed to remove file: path = "+d)
			errs = append(errs, err)
		}
		removedFiles = append(removedFiles, d)
	}
	if 0 < len(errs) {
		return nil, errors.New("failed to remove file")
	}

	return removedFiles, nil
}
