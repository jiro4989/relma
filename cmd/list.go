package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(commandList)
}

var commandList = &cobra.Command{
	Use:   "list",
	Short: "print installed GitHub Releases infomation",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := NewApp()
		if err != nil {
			return err
		}
		return a.CmdList(nil)
	},
}

func (a *App) CmdList(p *CommandLineListParam) error {
	rels, err := a.Config.ReadReleasesFile()
	if err != nil {
		return err
	}

	for _, rel := range rels {
		fmt.Println(rel.FormatVersion())
	}
	return nil
}
