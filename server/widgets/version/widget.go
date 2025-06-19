package version

import (
	"github.com/go-chi/chi/v5"
	"smartmirror.server/env"
	"smartmirror.server/widgets"
)

type VersionWidget struct{}

var _ widgets.Widget = (*VersionWidget)(nil)

func (v *VersionWidget) SetupEnv() {
	env.SetDefaultValue(envVersionHash, "notset")
}

func (v *VersionWidget) SetupRouter(router *chi.Mux) {
	router.HandleFunc("/version-hash", versionHashHandler)
}

func NewVersionWidget() *VersionWidget {
	return &VersionWidget{}
}
