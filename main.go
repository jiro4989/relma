package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/docopt/docopt-go"
)

type CommandLineParam struct {
	Command          string   `docopt:"command"`
	GitHubReleaseURL string   `docopt:"<github_release_url>"`
	Yes              bool     `docopt:"-y,--yes"`
	Args             []string `docopt:"<args>"`
}

const version = "v1.0.0"

const usage = `relma manages GitHub Releases versioning.

Usage:
  relma [command] [options]
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
  update  [-y | --yes]
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
	os.Exit(Main(os.Args[1:]))
}

func Main(args []string) int {
	opts, err := docopt.ParseArgs(usage, args, version)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	var clp CommandLineParam
	opts.Bind(&clp)

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	a := App{
		Config: Config{
			RelmaRoot: filepath.Join(home, "relma"),
		},
	}
	switch clp.Command {
	case "install":
		err = a.CmdInstall(clp.GitHubReleaseURL)
	case "update":
		p := CmdUpdateParam{
			Yes: clp.Yes,
		}
		err = a.CmdUpdate(&p)
	case "upgrade":
	}
	if err != nil {
		panic(err)
	}

	return 0
}
