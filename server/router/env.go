package router

import (
	"os"

	"smartmirror.server/env"
)

const (
	envCorsAllowedOrigin = "CORS_ALLOWED_ORIGIN"
)

func setupEnv() {
	env.ValidateEnvKeys([]string{
		envCorsAllowedOrigin,
	})
}

func getCorsAllowedOrigin() string {
	return os.Getenv(envCorsAllowedOrigin)
}
