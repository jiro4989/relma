package main

import (
	"log"
	"os"
)

var (
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
)

func Error(msgs ...interface{}) {
	errorLogger.Println(msgs...)
}
