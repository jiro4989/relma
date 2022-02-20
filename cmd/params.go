package cmd

type CommandLineParam struct {
	Command string   `docopt:"<command>"`
	Args    []string `docopt:"<args>"`
}

type CommandLineInitParam struct {
	Init bool
}

type CommandLineEditParam struct {
	Edit   bool
	Editor string `docopt:"-e,--editor"`
}

type CommandLineInstallParam struct {
	Install          bool
	GitHubReleaseURL string `docopt:"<github_release_url>"`
	File             string `docopt:"-f,--file"`
}

type CommandLineUpdateParam struct {
	Update   bool
	Releases []string `docopt:"<releases>"`
}

type CommandLineUpgradeParam struct {
	Upgrade   bool
	Yes       bool   `docopt:"-y,--yes"`
	OwnerRepo string `docopt:"<owner/repo>"`
}

type CommandLineUninstallParam struct {
	Uninstall bool
	OwnerRepo string `docopt:"<owner/repo>"`
}

type CommandLineListParam struct {
	List bool
}

type CommandLineRootParam struct {
	Root bool
}
