package kvg

import "os"

const (
	envHomeStopID = "KVG_HOME_STOP_ID"
)

func getEnvKeys() []string {
	return []string{
		envHomeStopID,
	}
}

func getHomeStopID() string {
	return os.Getenv(envHomeStopID)
}
