package version

import (
	"github.com/go-chi/chi/v5"
	"smartmirror.server/env"
)

type VersionWidget struct{}

func (v *VersionWidget) SetupEnv() {
	env.SetDefaultValue(envVersionHash, "notset")
	env.ValidateEnvKeys([]string{
		envVersionHash,
	})
}

func (v *VersionWidget) SetupRouter(router *chi.Mux) {
	router.HandleFunc("/version-hash", versionHashHandler)
}

func NewVersionWidget() *VersionWidget {
	return &VersionWidget{}
}
