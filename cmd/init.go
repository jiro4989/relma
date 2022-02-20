package main

import (
	"fmt"
	"os"
)

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
