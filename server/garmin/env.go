package garmin

import "os"

const (
	envUsername = "GARMIN_USERNAME"
	envPassword = "GARMIN_PASSWORD"
)

func GetEnvKeys() []string {
	return []string{
		envUsername,
		envPassword,
	}
}

func SetDefaultEnv() {

}

func getUsername() string {
	return os.Getenv(envUsername)
}

func getPassword() string {
	return os.Getenv(envPassword)
}
