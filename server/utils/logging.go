package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/getsentry/sentry-go"
)

type Logger struct {
	Info func(formattedString string, args ...interface{})
}

func NewLogger(widget string) Logger {
	return Logger{
		Info: func(formattedString string, args ...interface{}) {
			logLine := fmt.Sprintf("[%s]: %s", widget, fmt.Sprintf(formattedString, args...))

			log.Println(logLine)

			ctx := context.Background()
			sentryLogger := sentry.NewLogger(ctx)
			sentryLogger.Info().Emit(logLine)
		},
	}
}
