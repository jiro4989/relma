package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "r",
	Short: "",
	Long: ``,
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
		return a.CmdRoot(nil)
	},
}

func (a *App) CmdRoot(p *CommandLineRootParam) error {
	fmt.Println(a.Config.RelmaRoot)
	return nil
}
