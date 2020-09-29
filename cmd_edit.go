package main

import (
	"os"
	"os/exec"
)

func (a *App) CmdEdit(p *CommandLineEditParam) error {
	f, err := ConfigFile()
	if err != nil {
		return err
	}

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
