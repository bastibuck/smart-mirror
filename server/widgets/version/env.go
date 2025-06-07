package version

import "os"

const (
	envVersionHash = "VERSION_HASH"
)

func getVersionHash() string {
	return os.Getenv(envVersionHash)
}
