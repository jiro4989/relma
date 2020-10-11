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

func NewApp() (App, error) {
	conf, err := ReadConfigFile()
	if err != nil {
		return App{}, err
	}
	app := App{
		Config: conf,
	}
	return app, nil
}

func (a *App) SaveReleases(rels Releases) error {
	b, err := json.MarshalIndent(&rels, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(a.Config.ReleasesFile(), b, os.ModePerm); err != nil {
		return err
	}

	return nil
}

type Config struct {
	RelmaRoot string
}

func DefaultConfig() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	c := Config{
		RelmaRoot: filepath.Join(home, appName),
	}
	return c, nil
}

func ConfigDir() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	p := filepath.Join(dir, appName)
	return p, nil
}

func CreateConfigDir() (string, error) {
	dir, err := ConfigDir()
	if err != nil {
		return "", err
	}

	_, err = os.Stat(dir)
	if !os.IsNotExist(err) {
		return dir, nil
	}

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}
	return dir, nil
}

func ConfigFile() (string, error) {
	dir, err := ConfigDir()
	if err != nil {
		return "", err
	}

	p := filepath.Join(dir, "config.json")
	return p, nil
}

func ReadConfigFile() (Config, error) {
	file, err := ConfigFile()
	if err != nil {
		return Config{}, err
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return Config{}, err
	}

	var conf Config
	err = json.Unmarshal(b, &conf)
	if err != nil {
		return Config{}, err
	}
	return conf, nil
}

func CreateConfigFile(c Config) (string, error) {
	file, err := ConfigFile()
	if err != nil {
		return "", err
	}

	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(file, b, os.ModePerm)
	if err != nil {
		return "", err
	}
	return file, nil
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
	return ReadReleasesFile(c.ReleasesFile())
}

func ReadReleasesFile(path string) (Releases, error) {
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
