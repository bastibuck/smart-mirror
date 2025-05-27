package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	// required env vars
	envFrontendUrl = "FRONTEND_URL"

	envStravaClientId     = "STRAVA_CLIENT_ID"
	envStravaClientSecret = "STRAVA_CLIENT_SECRET"

	envCorsAllowedOrigin = "CORS_ALLOWED_ORIGIN"

	// optional env vars
	envServerPort  = "SERVER_PORT"
	envVersionHash = "VERSION_HASH"
)

var requiredEnvKeys = []string{
	envFrontendUrl,
	envStravaClientId,
	envStravaClientSecret,
	envCorsAllowedOrigin,
	envServerPort,
	envVersionHash,
}

func SetupEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	if os.Getenv(envVersionHash) == "" {
		os.Setenv(envVersionHash, "notset")
	}

	if os.Getenv(envServerPort) == "" {
		os.Setenv(envServerPort, "8080")
	}

	var anyMissingEnv bool
	for _, key := range requiredEnvKeys {
		if os.Getenv(key) == "" {
			anyMissingEnv = true
			fmt.Printf("Error: missing environment variable: %s\n", key)
		}

	}

	if anyMissingEnv {
		os.Exit(1)
	}
}

// getters

func GetVersionHash() string {
	return os.Getenv(envVersionHash)
}

func GetFrontendUrl() string {
	return os.Getenv(envFrontendUrl)
}

func GetStravaClientId() string {
	return os.Getenv(envStravaClientId)
}

func GetStravaClientSecret() string {
	return os.Getenv(envStravaClientSecret)
}

func GetServerPort() string {
	return os.Getenv(envServerPort)
}

func GetCorsAllowedOrigin() string {
	return os.Getenv(envCorsAllowedOrigin)
}
