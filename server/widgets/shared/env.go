package shared

import (
	"fmt"
	"os"

	"smartmirror.server/env"
)

const (
	envAppMode      = "APP_MODE" // "development" or "production"
	envGrafanaToken = "GRAFANA_TOKEN"
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

	// Warn if GRAFANA_TOKEN is not set
	if os.Getenv(envGrafanaToken) == "" {
		fmt.Println("Warning: GRAFANA_TOKEN not set, logging to Grafana will fail")
	}
}

func GetAppMode() string {
	return os.Getenv(envAppMode)
}

func GetGrafanaToken() string {
	return os.Getenv(envGrafanaToken)
}
