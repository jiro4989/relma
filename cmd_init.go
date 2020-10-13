package main

import (
	"os"
)

func (a *App) CmdInit() error {
	_, err := CreateConfigDir()
	if err != nil {
		return err
	}

	conf, err := DefaultConfig()
	if err != nil {
		return err
	}

	_, err = CreateConfigFile(conf)
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

	Message("initialize successful")

	return nil
}
