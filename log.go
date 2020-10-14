package main

import (
	"fmt"
	"log"
	"os"
)

var (
	msgLogger   = log.New(os.Stdout, "", 0)
	errorLogger = log.New(os.Stderr, "", 0)
)

func Message(msgs ...interface{}) {
	msgLogger.Println(msgs...)
}

func messageSkel(l *log.Logger, ok string, msgs ...string) {
	s := fmt.Sprintf("%-8s", ok)
	s += fmt.Sprintf("%-16s", msgs[0])
	if 1 < len(msgs) {
		s += fmt.Sprintf("%-16s", msgs[1])
	}
	l.Println(s)
}

func MessageOK(msgs ...string) {
	messageSkel(msgLogger, "ok", msgs...)
}

func MessageNG(msgs ...string) {
	messageSkel(errorLogger, "ng", msgs...)
}

func Error(msgs ...interface{}) {
	errorLogger.Println(msgs...)
}
