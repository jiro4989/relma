package cmd

type CommandLineParam struct {
	Command string   `docopt:"<command>"`
	Args    []string `docopt:"<args>"`
}
