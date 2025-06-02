package shared

import (
	"fmt"
	"os"
)

const (
	envCorsAllowedOrigin = "CORS_ALLOWED_ORIGIN"
	envServerPort        = "SERVER_PORT"
	envAppMode           = "APP_MODE" // "development" or "production"
)

func GetEnvKeys() []string {
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

func SetDefaultEnv() {
	if os.Getenv(envServerPort) == "" {
		os.Setenv(envServerPort, "8080")
	}

	if os.Getenv(envAppMode) == "" || !allowedModeValuesMap[os.Getenv(envAppMode)] {
		fmt.Println("Warning: APP_MODE not set or invalid, defaulting to 'production'")

		os.Setenv(envAppMode, "production")
	}
}

func GetServerPort() string {
	return os.Getenv(envServerPort)
}

func GetCorsAllowedOrigin() string {
	return os.Getenv(envCorsAllowedOrigin)
}

func GetAppMode() string {
	return os.Getenv(envAppMode)
}
