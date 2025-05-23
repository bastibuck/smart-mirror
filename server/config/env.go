package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	// required env vars
	EnvStravaAccessToken = "STRAVA_ACCESS_TOKEN"
	EnvStravaAthleteID   = "STRAVA_ATHLETE_ID"
	EnvCorsAllowedOrigin = "CORS_ALLOWED_ORIGIN"

	// optional env vars
	EnvServerPort  = "SERVER_PORT"
	EnvVersionHash = "VERSION_HASH"
)

var RequiredEnvKeys = []string{
	EnvStravaAccessToken,
	EnvStravaAthleteID,
	EnvCorsAllowedOrigin,
	EnvServerPort,
	EnvVersionHash,
}

func SetAndValidateEnvVars() {
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
