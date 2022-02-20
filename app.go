package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jiro4989/relma/downloader"
	"github.com/jiro4989/relma/thirdparty/github"
)

type App struct {
	Config        Config
	UserHomeDir   string
	UserConfigDir string
	GitHubClient  github.GitHubClientInterface
	Downloader    downloader.DownloaderInterface
}

func NewApp() (App, error) {
	var app App
	if err := app.SetUserEnv(); err != nil {
		return App{}, err
	}

	conf, err := app.ReadConfigFile()
	if err != nil {
		return App{}, err
	}

	app.Config = conf
	app.GitHubClient = github.NewClient()
	app.Downloader = downloader.NewDownloader()
	return app, nil
}

func (a *App) SetUserEnv() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	a.UserHomeDir = home

	conf, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	a.UserConfigDir = conf
	return nil
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

func (a *App) DefaultConfig() Config {
	home := a.UserHomeDir
	c := Config{
		RelmaRoot: filepath.Join(home, appName),
	}
	return c
}

func (a *App) ConfigDir() string {
	dir := a.UserConfigDir
	p := filepath.Join(dir, appName)
	return p
}

func (a *App) CreateConfigDir() (string, error) {
	dir := a.ConfigDir()

	_, err := os.Stat(dir)
	if !os.IsNotExist(err) {
		return dir, nil
	}

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}
	return dir, nil
}

func (a *App) ConfigFile() string {
	dir := a.ConfigDir()
	p := filepath.Join(dir, "config.json")
	return p
}

func (a *App) ReadConfigFile() (Config, error) {
	file := a.ConfigFile()

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

func (a *App) CreateConfigFile(c Config) (string, error) {
	file := a.ConfigFile()

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
