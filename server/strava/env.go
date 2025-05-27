package strava

import "os"

const (
	envStravaClientId        = "STRAVA_CLIENT_ID"
	envStravaClientSecret    = "STRAVA_CLIENT_SECRET"
	envStravaLoginSuccessUrl = "STRAVA_LOGIN_SUCCESS_URL"
	envStravaLoginFailureUrl = "STRAVA_LOGIN_FAILURE_URL"
)

func GetEnvKeys() []string {
	return []string{
		envStravaClientId,
		envStravaClientSecret,
		envStravaLoginSuccessUrl,
		envStravaLoginFailureUrl,
	}
}

func SetDefaultEnv() {
	// not implemented, as Strava client ID and secret are required
}

func getStravaClientId() string {
	return os.Getenv(envStravaClientId)
}

func getStravaClientSecret() string {
	return os.Getenv(envStravaClientSecret)
}

func GetStravaLoginSuccessUrl() string {
	return os.Getenv(envStravaLoginSuccessUrl)
}

func GetStravaLoginFailureUrl() string {
	return os.Getenv(envStravaLoginFailureUrl)
}
