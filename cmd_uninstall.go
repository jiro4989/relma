package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func (a *App) CmdUninstall(owner, repo string) error {
	dir := a.Config.ReleasesDir()
	p := filepath.Join(dir, owner, repo)
	err := os.RemoveAll(p)
	if err != nil {
		return err
	}

	fmt.Println("uninstall successful")

	return nil
}
