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

func MessageSkel(msgs ...string) {
	s := fmt.Sprintf("%-8s", msgs[0])
	if 1 < len(msgs) {
		s += fmt.Sprintf("%-16s", msgs[1])
	}
	if 2 < len(msgs) {
		s += fmt.Sprintf("%-16s", msgs[2])
	}
	msgLogger.Println(s)
}

func MessageOK(msgs ...string) {
	m := append([]string{"ok"}, msgs...)
	MessageSkel(m...)
}

func MessageEmpty(msgs ...string) {
	m := append([]string{" "}, msgs...)
	MessageSkel(m...)
}

func Error(msgs ...interface{}) {
	errorLogger.Println(msgs...)
}
