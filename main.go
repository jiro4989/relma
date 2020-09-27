package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/docopt/docopt-go"
)

type CommandLineParam struct {
	Command string   `docopt:"<command>"`
	Args    []string `docopt:"<args>"`
}

type CommandLineInstallParam struct {
	Install          bool
	GitHubReleaseURL string `docopt:"<github_release_url>"`
}

type CommandLineUpdateParam struct {
	Yes      bool     `docopt:"-y,--yes"`
	Releases []string `docopt:"<releases>"`
}

type CommandLineUpgradeParam struct {
	Command string   `docopt:"command"`
	Args    []string `docopt:"<args>"`
	Yes     bool     `docopt:"-y,--yes"`
}

const (
	version = "v1.0.0"
	usage   = `relma manages GitHub Releases versioning.

usage:
  relma [options] <command> [<args>...]
  relma -h | --help
  relma --version

commands:
  init         initialize config file.
  edit         edit config file.
  install      install GitHub Releases.
  list
  update       update installed version infomation.
  upgrade      upgrade installed GitHub Releases.
  uninstall    uninstall GitHub Releases.

options:
  -h, --help    print this help
  --version     print version
`

	usageInstall = `usage: relma install [options] <github_release_url>

options:
  -h, --help       print this help
`

	usageUpdate = `usage: relma update [options] [<releases>...]

options:
  -h, --help       print this help
  -y, --yes        yes
`
)

func main() {
	os.Exit(Main(os.Args[1:]))
}

func Main(args []string) int {
	parser := &docopt.Parser{OptionsFirst: true}

	opts, err := parser.ParseArgs(usage, args, version)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	var clp CommandLineParam
	err = opts.Bind(&clp)
	if err != nil {
		panic(err)
	}

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
		args := []string{clp.Command}
		args = append(args, opts["<args>"].([]string)...)
		opts, err := docopt.ParseArgs(usageInstall, args, "")
		if err != nil {
			panic(err)
		}
		var clp CommandLineInstallParam
		err = opts.Bind(&clp)
		if err != nil {
			panic(err)
		}

		err = a.CmdInstall(clp.GitHubReleaseURL)
	case "update":
		args := []string{clp.Command}
		args = append(args, opts["<args>"].([]string)...)
		opts, err := docopt.ParseArgs(usageUpdate, args, "")
		if err != nil {
			panic(err)
		}
		var clp CommandLineUpdateParam
		opts.Bind(&clp)

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
