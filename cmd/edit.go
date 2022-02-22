package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	commandEdit.Flags().StringVarP(&commandLineEditParam.Editor, "editor", "e", "", "edit config file")

	rootCmd.AddCommand(commandEdit)
}

type CommandLineEditParam struct {
	Editor string
}

var commandLineEditParam CommandLineEditParam

var commandEdit = &cobra.Command{
	Use:   "edit",
	Short: "Edit config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := NewApp()
		if err != nil {
			return err
		}
		return a.CmdEdit(&commandLineEditParam)
	},
}

func (a *App) CmdEdit(p *CommandLineEditParam) error {
	f := a.ConfigFile()

	editor := os.Getenv("EDITOR")
	if p.Editor != "" {
		editor = p.Editor
	}

	cmd := exec.Command(editor, f)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
