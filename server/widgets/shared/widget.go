package shared

import "github.com/go-chi/chi/v5"

type SharedWidget struct{}

func (v *SharedWidget) SetupEnv() {
	setDefaultEnv()
}

func (v *SharedWidget) SetupRouter(router *chi.Mux) {
	// no routes to setup
}

func NewSharedWidget() *SharedWidget {
	return &SharedWidget{}
}
