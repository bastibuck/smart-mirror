package garmin

import (
	"github.com/go-chi/chi/v5"
	"smartmirror.server/env"
)

type GarminWidget struct{}

func (v *GarminWidget) SetupEnv() {
	env.ValidateEnvKeys(getEnvKeys())
}

func (v *GarminWidget) SetupRouter(router *chi.Mux) {
	router.HandleFunc("/steps/today", stepsTodayHandler)
	router.HandleFunc("/steps/weekly", stepsThisWeekHandler)
}

func NewGarminWidget() *GarminWidget {
	return &GarminWidget{}
}
