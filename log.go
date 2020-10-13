package main

import (
	"log"
	"os"
)

var (
	msgLogger   = log.New(os.Stdout, "", 0)
	errorLogger = log.New(os.Stdout, appName+": ERROR ", log.Ldate|log.Ltime)
)

func Message(msgs ...interface{}) {
	msgLogger.Println(msgs...)
}

func Error(msgs ...interface{}) {
	errorLogger.Println(msgs...)
}
