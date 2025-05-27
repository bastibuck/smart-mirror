package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"smartmirror.server/strava"
	"smartmirror.server/version"
)

const (
	envCorsAllowedOrigin = "CORS_ALLOWED_ORIGIN"
	envServerPort        = "SERVER_PORT"
)

var requiredEnvKeys = []string{
	envCorsAllowedOrigin,
	envServerPort,
}

func SetupEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	// defaults
	if os.Getenv(envServerPort) == "" {
		os.Setenv(envServerPort, "8080")
	}

	version.SetDefaultEnv()
	strava.SetDefaultEnv()

	// validate required environment variables
	var anyMissingEnv bool

	var allEnvKeys []string = requiredEnvKeys
	allEnvKeys = append(allEnvKeys, strava.GetEnvKeys()...)
	allEnvKeys = append(allEnvKeys, version.GetEnvKeys()...)

	for _, key := range allEnvKeys {
		if os.Getenv(key) == "" {
			anyMissingEnv = true
			fmt.Printf("Error: missing environment variable: %s\n", key)
		}
	}

	if anyMissingEnv {
		os.Exit(1)
	}
}

func GetServerPort() string {
	return os.Getenv(envServerPort)
}

func GetCorsAllowedOrigin() string {
	return os.Getenv(envCorsAllowedOrigin)
}
