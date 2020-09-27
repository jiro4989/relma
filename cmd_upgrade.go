package main

import (
	"errors"
	"strings"
	"time"
)

type CmdUpgradeParam struct {
	Yes bool
}

func (a *App) CmdUpgrade(p *CmdUpgradeParam) error {
	rels, err := a.Config.ReadReleasesFile()
	if err != nil {
		return err
	}
	if rels == nil {
		return errors.New("installed releases don't exist")
	}
	for _, rel := range rels {
		if rel.Version == rel.LatestVersion {
			continue
		}
		url := strings.ReplaceAll(rel.URL, rel.Version, rel.LatestVersion)
		err := a.CmdInstall(url)
		if err != nil {
			return err
		}

		time.Sleep(1 * time.Second)
	}
	return nil
}
