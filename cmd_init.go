package main

import (
	"os"
)

func (a *App) CmdInit() error {
	configDir, err := CreateConfigDir()
	if err != nil {
		return err
	}
	Info("created " + configDir)

	conf, err := DefaultConfig()
	if err != nil {
		return err
	}

	confFile, err := CreateConfigFile(conf)
	if err != nil {
		return err
	}
	Info("created " + confFile)

	paths := []string{
		conf.RelmaRoot,
		conf.BinDir(),
		conf.ReleasesDir(),
	}
	for _, path := range paths {
		_, err := os.Stat(path)
		if !os.IsNotExist(err) {
			Info(path + " was already existed")
			continue
		}

		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
		Info("created " + path)
	}

	Info("initialize successful")

	return nil
}
