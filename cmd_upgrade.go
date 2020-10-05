package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type CmdUpgradeParam struct {
	OwnerRepo string
	Yes       bool
}

func (a *App) CmdUpgrade(p *CmdUpgradeParam) error {
	rels, err := a.Config.ReadReleasesFile()
	if err != nil {
		return err
	}

	rels, err = searchReleaseOrDefault(rels, p.OwnerRepo)
	if err != nil {
		return err
	}

	targets := upgradableReleases(rels)
	if len(targets) < 1 {
		fmt.Println("upgradable releases were not existed")
		return nil
	}

	fmt.Print("update? [y/n] > ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	if strings.ToLower(sc.Text()) != "y" {
		fmt.Println("not upgrade")
		return nil
	}

	for _, rel := range targets {
		url := strings.ReplaceAll(rel.URL, rel.Version, rel.LatestVersion)
		err := a.CmdInstall(url)
		if err != nil {
			return err
		}

		time.Sleep(1 * time.Second)
	}
	fmt.Println("upgrade successful")

	return nil
}

func searchReleaseOrDefault(rels Releases, ownerRepo string) (Releases, error) {
	if len(rels) < 1 {
		return nil, errors.New("installed releases don't exist")
	}

	if ownerRepo != "" {
		var err error
		rels, err = searchRelease(rels, ownerRepo)
		if err != nil {
			return nil, err
		}
		if len(rels) < 1 {
			msg := fmt.Sprintf("%s was not installed", ownerRepo)
			return nil, errors.New(msg)
		}
	}

	return rels, nil
}

func searchRelease(rels Releases, ownerRepo string) (Releases, error) {
	var retRels Releases
	for _, rel := range rels {
		if ok, err := rel.EqualRepo(ownerRepo); err != nil {
			return nil, err
		} else if !ok {
			continue
		}
		retRels = append(retRels, rel)
		return retRels, nil
	}
	return nil, nil
}

func upgradableReleases(rels Releases) Releases {
	var upgradables Releases
	for _, rel := range rels {
		if rel.Version == rel.LatestVersion {
			continue
		}
		upgradables = append(upgradables, rel)
		info := fmt.Sprintf("%s/%s %s -> %s", rel.Owner, rel.Repo, rel.Version, rel.LatestVersion)
		fmt.Println(info)
	}
	return upgradables
}
