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

	return Logger{
		Info: func(formattedString string, args ...interface{}) {
			timestamp := time.Now().Format("2006-01-02 15:04:05")

			log.Printf("[%s] %s: %s\n", widget, timestamp, fmt.Sprintf(formattedString, args...))
		},
	}
}
