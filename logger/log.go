package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
)

func Error(msgs ...interface{}) {
	var sb strings.Builder
	for _, msg := range msgs {
		sb.WriteString(fmt.Sprintf("%v", msg))
	}
	s := sb.String()
	errorLogger.Output(2, s)
}
