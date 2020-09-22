package main

import (
	"path/filepath"
)

type App struct {
	Config Config
}

type Config struct {
	GhrPkgRoot string
}

func (c *Config) PackageDir() string {
	return filepath.Join(c.GhrPkgRoot, "pkg")
}

const usage = `ghrpkg manages GitHub Releases versioning.

Usage:
  ghrpkg [commands] [options]
  ghrpkg -h | --help
  ghrpkg --version

Examples:
  $ ghrpkg init

  $ ghrpkg edit

  $ ghrpkg install https://github.com/jiro4989/nimjson/releases/download/v1.2.6/nimjson_linux.tar.gz

  $ ghrpkg list --upgradable

  $ ghrpkg show jiro4989/nimjson

  $ ghrpkg upgrade jiro4989/nimjson v1.2.7

  $ ghrpkg upgrade --all

  $ ghrpkg remove jiro4989/nimjson

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
