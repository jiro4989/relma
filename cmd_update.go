package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v32/github"
)

type CmdUpdateParam struct {
	Yes bool
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
		fmt.Println(rel.Owner+"/"+rel.Repo, "current_tag:", rel.Version, "available_latest_tag:", latestTag)
		rel.LatestVersion = latestTag
		time.Sleep(1 * time.Second)
	}

	fmt.Println("")

	doUpdate := p.Yes
	if !doUpdate {
		fmt.Print("update? [y/n] > ")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		if strings.ToLower(sc.Text()) == "y" {
			doUpdate = true
		}
	}

	if !doUpdate {
		fmt.Println("not updated")
		return nil
	}

	b, err := json.Marshal(rels)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(a.Config.ReleasesFile(), b, os.ModePerm)
	if err != nil {
		return err
	}

	fmt.Println("update successful")

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
