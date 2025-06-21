package shared

import (
	"github.com/go-chi/chi/v5"
	"smartmirror.server/widgets"
)

type SharedWidget struct{}

var _ widgets.Widget = (*SharedWidget)(nil)

func (v *SharedWidget) SetupEnv() {
	setDefaultEnv()
}

func (v *SharedWidget) SetupRouter(router *chi.Mux) {
	// no routes to setup
}

func NewSharedWidget() *SharedWidget {
	return &SharedWidget{}
}
