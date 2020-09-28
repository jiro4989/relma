package main

import (
	"fmt"
	"os"
)

func (a *App) CmdInit() error {
	configDir, err := CreateConfigDir()
	if err != nil {
		return err
	}
	fmt.Println("created " + configDir)

	conf, err := DefaultConfig()
	if err != nil {
		return err
	}

	confFile, err := CreateConfigFile(conf)
	if err != nil {
		return err
	}
	fmt.Println("created " + confFile)

	paths := []string{
		conf.RelmaRoot,
		conf.BinDir(),
		conf.ReleasesDir(),
	}
	for _, path := range paths {
		_, err := os.Stat(path)
		if !os.IsNotExist(err) {
			fmt.Println(path + " was already existed")
			continue
		}

		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Println("created " + path)
	}

	fmt.Println("initialize successful")

	return nil
}
