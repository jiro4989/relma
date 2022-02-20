package main

import (
	"errors"
	"fmt"
)

type CmdUpdateParam struct {
	Releases []string
}

func (a *App) CmdUpdate(p *CmdUpdateParam) error {
	rels, err := a.Config.ReadReleasesFile()
	if err != nil {
		return err
	}
	if rels == nil {
		return errors.New("installed releases don't exist")
	}

	for i := 0; i < len(rels); i++ {
		rel := &rels[i] // for override 'LatestVersion'

		latestTag, err := a.GitHubClient.FetchLatestTag(rel.Owner, rel.Repo)
		if err != nil {
			return err
		}
		if rel.Version != latestTag {
			fmt.Println(rel.FormatVersion() + " -> " + latestTag)
		} else {
			fmt.Println(rel.FormatSimpleInformation() + " -> same")
		}
		rel.LatestVersion = latestTag
		Sleep()
	}

	err = a.SaveReleases(rels)
	if err != nil {
		return err
	}

	fmt.Println("update successful")

	return nil
}
