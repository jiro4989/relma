package main

import (
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

func MessageOK(msgs ...interface{}) {
	msgs = append([]interface{}{"ok     "}, msgs...)
	msgLogger.Println(msgs...)
}

func Error(msgs ...interface{}) {
	errorLogger.Println(msgs...)
}
