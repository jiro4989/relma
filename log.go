package main

import (
	"fmt"
	"log"
	"os"
)

var (
	msgLogger   = log.New(os.Stdout, "", 0)
	errorLogger = log.New(os.Stderr, appName+": [ERROR] ", 0)
)

func Message(msgs ...interface{}) {
	msgLogger.Println(msgs...)
}

func MessageOK(msgs ...string) {
	s := fmt.Sprintf("%-8s", "ok")
	s += fmt.Sprintf("%-16s", msgs[0])
	if 1 < len(msgs) {
		s += fmt.Sprintf("%-16s", msgs[1])
	}
	msgLogger.Println(s)
}

func Error(msgs ...interface{}) {
	errorLogger.Println(msgs...)
}
