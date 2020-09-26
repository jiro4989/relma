package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type App struct {
	Config Config
}

type Config struct {
	RelmaRoot string
}

func (c *Config) ReleasesDir() string {
	return filepath.Join(c.RelmaRoot, "releases")
}

func (c *Config) ReleasesFile() string {
	return filepath.Join(c.RelmaRoot, "releases.json")
}

func (c *Config) BinDir() string {
	return filepath.Join(c.RelmaRoot, "bin")
}

func (c *Config) ReadReleasesFile() (Releases, error) {
	return readReleasesFile(c.ReleasesFile())
}

func readReleasesFile(path string) (Releases, error) {
	var rels Releases
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(b, &rels); err != nil {
			return nil, err
		}
	}
	return rels, nil
}
