package windspeed

import (
	"github.com/go-chi/chi/v5"
	"smartmirror.server/env"
	"smartmirror.server/widgets"
)

type WindspeedWidget struct{}

var _ widgets.Widget = (*WindspeedWidget)(nil)

func (w *WindspeedWidget) SetupEnv() {
	env.ValidateEnvKeys(getEnvKeys())

	// TODO? validate GPS coordinates
}

func (w *WindspeedWidget) SetupRouter(router *chi.Mux) {
	router.HandleFunc("/windspeed", windspeedHandler)
}

func NewWindspeedWidget() *WindspeedWidget {
	return &WindspeedWidget{}
}
