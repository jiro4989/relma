package main

import (
	"fmt"
)

func (a *App) CmdList(p *CommandLineListParam) error {
	rels, err := a.Config.ReadReleasesFile()
	if err != nil {
		return err
	}

	for _, rel := range rels {
		s := fmt.Sprintf("%s/%s %s", rel.Owner, rel.Repo, rel.Version)
		fmt.Println(s)
	}
	return nil
}
