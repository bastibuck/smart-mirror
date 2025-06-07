package app

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"smartmirror.server/env"
)

const (
	envCorsAllowedOrigin = "CORS_ALLOWED_ORIGIN"
	envServerPort        = "SERVER_PORT"
	envAppMode           = "APP_MODE" // "development" or "production"
)

func getEnvKeys() []string {
	return []string{
		envCorsAllowedOrigin,
		envServerPort,
		envAppMode,
	}
}

var allowedModeValuesMap = map[string]bool{
	"development": true,
	"production":  true,
}

func setDefaultEnv() {
	env.SetDefaultValue(envServerPort, "8080")

	if os.Getenv(envAppMode) == "" || !allowedModeValuesMap[os.Getenv(envAppMode)] {
		fmt.Println("Warning: APP_MODE not set or invalid, defaulting to 'production'")

		os.Setenv(envAppMode, "production")
	}
}

func getServerPort() string {
	return os.Getenv(envServerPort)
}

func getCorsAllowedOrigin() string {
	return os.Getenv(envCorsAllowedOrigin)
}

func GetAppMode() string {
	return os.Getenv(envAppMode)
}

func setupAppEnv() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	setDefaultEnv()

	env.ValidateEnvKeys(getEnvKeys())
}
