package main

import (
	"os"

	"github.com/jiro4989/relma/cmd"
	"github.com/jiro4989/relma/logger"
)

var (
	version = "dev"
)

func main() {
	cmd.Version = version
	if err := cmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
