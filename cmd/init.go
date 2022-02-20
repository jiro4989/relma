package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(commandInit)
}

var commandInit = &cobra.Command{
	Use:   "init",
	Short: "print installed GitHub Releases infomation",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := NewApp()
		if err != nil {
			return err
		}
		return a.CmdInit()
	},
}

func (a *App) CmdInit() error {
	_, err := a.CreateConfigDir()
	if err != nil {
		return err
	}

	conf := a.DefaultConfig()

	_, err = a.CreateConfigFile(conf)
	if err != nil {
		return err
	}

	paths := []string{
		conf.RelmaRoot,
		conf.BinDir(),
		conf.ReleasesDir(),
	}
	for _, path := range paths {
		_, err := os.Stat(path)
		if !os.IsNotExist(err) {
			continue
		}

		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	fmt.Println("initialize successful")

	return nil
}
