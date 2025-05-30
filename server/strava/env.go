package strava

import (
	"os"
	"strconv"
)

const (
	envStravaClientId        = "STRAVA_CLIENT_ID"
	envStravaClientSecret    = "STRAVA_CLIENT_SECRET"
	envStravaLoginSuccessUrl = "STRAVA_LOGIN_SUCCESS_URL"
	envStravaLoginFailureUrl = "STRAVA_LOGIN_FAILURE_URL"

	// for local dev
	envStravaAccessTokenOverride  = "STRAVA_ACCESS_TOKEN_OVERRIDE"
	envStravaRefreshTokenOverride = "STRAVA_REFRESH_TOKEN_OVERRIDE"
	envStravaAthleteIdOverride    = "STRAVA_ATHLETE_ID_OVERRIDE"
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
	// This is used for local development to override the access
	// so that you don't have to go through the OAuth flow every time.
	GLOBAL_StravaAccessToken = os.Getenv(envStravaAccessTokenOverride)
	GLOBAL_StravaRefreshToken = os.Getenv(envStravaRefreshTokenOverride)
	GLOBAL_StravaAthleteId = func() int {
		v := os.Getenv(envStravaAthleteIdOverride)
		i, _ := strconv.Atoi(v)
		return i
	}()
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
