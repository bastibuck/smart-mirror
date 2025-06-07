package env

import (
	"fmt"
	"os"
)

func ValidateEnvKeys(envKeys []string) {
	var anyMissingEnv bool

	for _, key := range envKeys {
		if os.Getenv(key) == "" {
			anyMissingEnv = true
			fmt.Printf("Error: missing environment variable: %s\n", key)
		}
	}

	if anyMissingEnv {
		os.Exit(1)
	}
}

func SetDefaultValue(key, value string) {
	if os.Getenv(key) == "" {
		os.Setenv(key, value)
	}
}
