package shared

import "github.com/go-chi/chi/v5"

type SharedWidget struct{}

func (v *SharedWidget) SetupEnv() {
	setDefaultEnv()
}

func (v *SharedWidget) SetupRouter(router *chi.Mux) {}

func NewSharedWidget() *SharedWidget {
	return &SharedWidget{}
}
