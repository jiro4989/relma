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

	var targets Releases
	for _, rel := range rels {
		if rel.Version == rel.LatestVersion {
			continue
		}
		targets = append(targets, rel)
		info := fmt.Sprintf("%s/%s %s -> %s", rel.Owner, rel.Repo, rel.Version, rel.LatestVersion)
		fmt.Println(info)
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
