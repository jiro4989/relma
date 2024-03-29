package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jiro4989/relma/releases"
	"github.com/spf13/cobra"
)

type CommandLineUninstallParam struct {
	OwnerRepo string
}

func init() {
	rootCmd.AddCommand(commandUninstall)
}

var commandLineUninstallParam CommandLineUninstallParam

var commandUninstall = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall GitHub Releases",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := NewApp()
		if err != nil {
			return err
		}
		c := CmdUninstallParam{
			OwnerRepo: args[0],
		}
		return a.CmdUninstall(&c)
	},
}

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

	rels = releases.RemoveRelease(rels, rel)
	err = a.SaveReleases(rels)
	if err != nil {
		return err
	}

	fmt.Println(rel.FormatVersion())

	return nil
}

func uninstallableRelease(rels releases.Releases, ownerRepo string) (*releases.Release, error) {
	var rel *releases.Release
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

func (a *App) uninstallRelease(rel *releases.Release, p *CmdUninstallParam) ([]string, error) {
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
			Error("failed to remove file: path = " + d)
			errs = append(errs, err)
		}
		removedFiles = append(removedFiles, d)
	}
	if 0 < len(errs) {
		return nil, errors.New("failed to remove file")
	}

	return removedFiles, nil
}
