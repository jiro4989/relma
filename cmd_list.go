package main

func (a *App) CmdList(p *CommandLineListParam) error {
	rels, err := a.Config.ReadReleasesFile()
	if err != nil {
		return err
	}

	for _, rel := range rels {
		Message(rel.FormatSimpleInformation())
	}
	return nil
}
