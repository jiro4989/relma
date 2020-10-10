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
		fmt.Println(rel.FormatSimpleInformation())
	}
	return nil
}
