package main

import (
	"context"
	"errors"

	"github.com/google/go-github/v32/github"
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

	c := github.NewClient(nil)
	for i := 0; i < len(rels); i++ {
		rel := &rels[i] // for override 'LatestVersion'

		latestTag, err := fetchLatestTag(c, rel.Owner, rel.Repo)
		if err != nil {
			return err
		}
		if rel.Version != latestTag {
			MessageOK("update", rel.FormatVersion()+" -> "+latestTag)
		} else {
			MessageOK("update", rel.FormatSimpleInformation()+" -> same")
		}
		rel.LatestVersion = latestTag
		Sleep()
	}

	err = a.SaveReleases(rels)
	if err != nil {
		return err
	}

	MessageOK("update")

	return nil
}

func fetchLatestTag(c *github.Client, owner, repo string) (string, error) {
	rel, _, err := c.Repositories.ListReleases(context.Background(), owner, repo, nil)
	if err != nil {
		return "", err
	}
	var latestTag string
	for _, rel := range rel {
		r := *rel
		latestTag = r.GetTagName()
		return latestTag, nil
	}
	return "", nil
}
