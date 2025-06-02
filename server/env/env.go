package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"smartmirror.server/shared"
	"smartmirror.server/strava"
	"smartmirror.server/version"
)

func SetupEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	// set default environment variables
	shared.SetDefaultEnv()
	version.SetDefaultEnv()
	strava.SetDefaultEnv()

	// validate required environment variables
	var anyMissingEnv bool

	var allEnvKeys []string = []string{}

	allEnvKeys = append(allEnvKeys, shared.GetEnvKeys()...)
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
