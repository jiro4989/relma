package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/go-github/v32/github"
)

type CmdUpdateParam struct {
}

func (a *App) CmdUpdate(p *CmdUpdateParam) error {
	rels, err := a.Config.ReadReleasesFile()
	if err != nil {
		return err
	}
	if rels == nil {
		return errors.New("installed releases don't exist")
	}

	c := github.NewClient(nil)
	for _, rel := range rels {
		ghRels, _, err := c.Repositories.ListReleases(context.Background(), rel.Owner, rel.Repo, nil)
		if err != nil {
			return err
		}
		var latestTag string
		for _, ghRel := range ghRels {
			r := *ghRel
			latestTag = r.GetTagName()
			break
		}
		fmt.Println(rel.Owner+"/"+rel.Repo, "current_tag:", rel.Version, "available_latest_tag:", latestTag)
		time.Sleep(1 * time.Second)
	}
	return nil
}
