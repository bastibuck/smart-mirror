package app

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"smartmirror.server/env"
)

const (
	envServerPort = "SERVER_PORT"
	envSentryDsn  = "SENTRY_DSN"
)

func setupAppEnv() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	env.SetDefaultValue(envServerPort, "8080")
}

func getServerPort() string {
	return os.Getenv(envServerPort)
}

func getSentryDsn() string {
	return os.Getenv(envSentryDsn)
}
