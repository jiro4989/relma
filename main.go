package main

import (
	"fmt"
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

func (c *Config) BinDir() string {
	return filepath.Join(c.RelmaRoot, "bin")
}

const usage = `relma manages GitHub Releases versioning.

Usage:
  relma [commands] [options]
  relma -h | --help
  relma --version

Examples:
  $ relma init

  $ relma edit

  $ relma install https://github.com/jiro4989/nimjson/releases/download/v1.2.6/nimjson_linux.tar.gz

  $ relma list --upgradable

  $ relma show jiro4989/nimjson

  $ relma upgrade jiro4989/nimjson v1.2.7

  $ relma upgrade --all

  $ relma remove jiro4989/nimjson

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
