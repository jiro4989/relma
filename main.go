package main

import (
	"github.com/jiro4989/relma/cmd"
)

var (
	version = "dev"
)

func main() {
	cmd.Version = version
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
