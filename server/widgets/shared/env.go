package shared

import (
	"fmt"
	"os"

	"smartmirror.server/env"
)

const (
	envAppMode = "APP_MODE" // "development" or "production"
)

var allowedModeValuesMap = map[string]bool{
	"development": true,
	"production":  true,
}

func setDefaultEnv() {
	if os.Getenv(envAppMode) == "" || !allowedModeValuesMap[os.Getenv(envAppMode)] {
		fmt.Println("Warning: APP_MODE not set or invalid, defaulting to 'production'")

		env.SetDefaultValue(envAppMode, "production")
	}
}

func GetAppMode() string {
	return os.Getenv(envAppMode)
}
