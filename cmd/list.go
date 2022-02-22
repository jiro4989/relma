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
	Short: "Print installed GitHub Releases infomation",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := NewApp()
		if err != nil {
			return err
		}
		return a.CmdList()
	},
}

func (a *App) CmdList() error {
	rels, err := a.Config.ReadReleasesFile()
	if err != nil {
		return err
	}

	for _, rel := range rels {
		fmt.Println(rel.FormatVersion())
	}
	return nil
}
