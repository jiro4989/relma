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
		Info(rel.FormatSimpleInformation())
	}
	return nil
}
