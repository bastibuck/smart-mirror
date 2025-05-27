package version

import "os"

const (
	envVersionHash = "VERSION_HASH"
)

func GetEnvKeys() []string {
	return []string{
		envVersionHash,
	}
}

func SetDefaultEnv() {
	if os.Getenv(envVersionHash) == "" {
		os.Setenv(envVersionHash, "notset")
	}
}

func getVersionHash() string {
	return os.Getenv(envVersionHash)
}
