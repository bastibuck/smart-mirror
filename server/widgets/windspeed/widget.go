package windspeed

import (
	"github.com/go-chi/chi/v5"
	"smartmirror.server/widgets"
)

type WindspeedWidget struct{}

var _ widgets.Widget = (*WindspeedWidget)(nil)

func (w *WindspeedWidget) SetupEnv() {
	// No environment variables to set up
}

func (w *WindspeedWidget) SetupRouter(router *chi.Mux) {
	router.HandleFunc("/windspeed", windspeedHandler)
}

func NewWindspeedWidget() *WindspeedWidget {
	return &WindspeedWidget{}
}
