package cmd

import (
	"fmt"
)

func (a *App) CmdRoot(p *CommandLineRootParam) error {
	fmt.Println(a.Config.RelmaRoot)
	return nil
}
