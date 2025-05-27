package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	// required env vars
	EnvFrontendUrl = "FRONTEND_URL"

	EnvStravaClientId     = "STRAVA_CLIENT_ID"
	EnvStravaClientSecret = "STRAVA_CLIENT_SECRET"

	EnvCorsAllowedOrigin = "CORS_ALLOWED_ORIGIN"

	// optional env vars
	EnvServerPort  = "SERVER_PORT"
	EnvVersionHash = "VERSION_HASH"
)

var RequiredEnvKeys = []string{
	EnvFrontendUrl,
	EnvStravaClientId,
	EnvStravaClientSecret,
	EnvCorsAllowedOrigin,
	EnvServerPort,
	EnvVersionHash,
}

func SetAndValidateEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	if os.Getenv(EnvVersionHash) == "" {
		os.Setenv(EnvVersionHash, "notset")
	}

	if os.Getenv(EnvServerPort) == "" {
		os.Setenv(EnvServerPort, "8080")
	}

	for _, key := range RequiredEnvKeys {
		if os.Getenv(key) == "" {
			fmt.Printf("Error: missing required environment variable: %s\n", key)
			os.Exit(1)
		}
	}
}
