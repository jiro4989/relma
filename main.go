package main

import (
	"os"

	"github.com/jiro4989/relma/cmd"
	"github.com/jiro4989/relma/logger"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
