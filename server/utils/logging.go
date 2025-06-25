package utils

import (
	"fmt"
	"log"
	"time"
)

type Logger struct {
	Info func(formattedString string, args ...interface{})
}

func NewLogger(widget string) Logger {
	ts := time.Now().Format("2006-01-02 15:04:05")

	return Logger{
		Info: func(formattedString string, args ...interface{}) {
			log.Printf("[%s] %s: %s\n", widget, ts, fmt.Sprintf(formattedString, args...))
		},
	}
}
