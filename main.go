package main

import (
	"fmt"
	"path/filepath"
)

type App struct {
	Config Config
}

type Config struct {
	RelmRoot string
}

func (c *Config) PackageDir() string {
	return filepath.Join(c.RelmRoot, "pkg")
}

func (c *Config) BinDir() string {
	return filepath.Join(c.RelmRoot, "bin")
}

const usage = `relm manages GitHub Releases versioning.

Usage:
  relm [commands] [options]
  relm -h | --help
  relm --version

Examples:
  $ relm init

  $ relm edit

  $ relm install https://github.com/jiro4989/nimjson/releases/download/v1.2.6/nimjson_linux.tar.gz

  $ relm list --upgradable

  $ relm show jiro4989/nimjson

  $ relm upgrade jiro4989/nimjson v1.2.7

  $ relm upgrade --all

  $ relm remove jiro4989/nimjson

Commands:
  init
  install [-d | --dry-run] <github_release_url>
  list    [-u | --upgradable]
  upgrade show <package>
  upgrade <package> <version>
          [--all]
  remove <package>
  edit

Options:
  -h, --help    print this help
      --version print version
`

func main() {
	fmt.Println("hello world")
}
