package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	EnvStravaAccessToken = "STRAVA_ACCESS_TOKEN"
	EnvStravaAthleteID   = "STRAVA_ATHLETE_ID"
)

var RequiredEnvKeys = []string{
	EnvStravaAccessToken,
	EnvStravaAthleteID,
}

func ValidateEnvVars() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	for _, key := range RequiredEnvKeys {
		if os.Getenv(key) == "" {
			fmt.Printf("Error: missing required environment variable: %s\n", key)
			os.Exit(1)
		}
	}
}
