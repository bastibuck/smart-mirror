package kptncook

import (
	"github.com/go-chi/chi/v5"
	"smartmirror.server/widgets"
)

type KptnCookWidget struct{}

var _ widgets.Widget = (*KptnCookWidget)(nil)

func (v *KptnCookWidget) SetupEnv() {
	// No environment variables needed for KptnCook widget
}

func (v *KptnCookWidget) SetupRouter(router *chi.Mux) {
	router.HandleFunc("/recipes/daily", dailyRecipesHandler)
}

func NewKptnCookWidget() *KptnCookWidget {
	return &KptnCookWidget{}
}
