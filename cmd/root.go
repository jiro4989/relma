package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "dev"

var rootCmd = &cobra.Command{
	Use:     "main",
	Short:   "relma manages GitHub Releases versioning",
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(commandRoot)
}

var commandRoot = &cobra.Command{
	Use:   "root",
	Short: "print relma root directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := NewApp()
		if err != nil {
			return err
		}
		return a.CmdRoot()
	},
}

func (a *App) CmdRoot() error {
	fmt.Println(a.Config.RelmaRoot)
	return nil
}
