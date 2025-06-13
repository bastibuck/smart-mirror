package garmin

import "os"

const (
	envEmail    = "GARMIN_EMAIL"
	envPassword = "GARMIN_PASSWORD"
)

func getEnvKeys() []string {
	return []string{
		envEmail,
		envPassword,
	}
}

func getEmail() string {
	return os.Getenv(envEmail)
}

func getPassword() string {
	return os.Getenv(envPassword)
}
