package main

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, appName+": INFO ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stdout, appName+": ERROR ", log.Ldate|log.Ltime)
)

func Info(msgs ...interface{}) {
	infoLogger.Println(msgs...)
}

func Error(msgs ...interface{}) {
	errorLogger.Println(msgs...)
}
