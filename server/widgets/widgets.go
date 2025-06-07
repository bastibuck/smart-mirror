package widgets

import "github.com/go-chi/chi/v5"

type Widget interface {
	SetupEnv()
	SetupRouter(router *chi.Mux)
}

func RegisterWidgets(widgets []Widget, router *chi.Mux) {
	for _, widget := range widgets {
		widget.SetupEnv()
		widget.SetupRouter(router)
	}
}
