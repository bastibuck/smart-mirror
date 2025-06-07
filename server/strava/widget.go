package strava

import (
	"github.com/go-chi/chi/v5"
	"smartmirror.server/env"
)

type StravaWidget struct{}

func (v *StravaWidget) SetupEnv() {
	env.ValidateEnvKeys(getEnvKeys())
	setDefaultEnv()
}

func (v *StravaWidget) SetupRouter(router *chi.Mux) {
	router.Route("/strava", func(subRouter chi.Router) {
		subRouter.Get("/exchange-token", exchangeTokenHandler)
		subRouter.Get("/creds", credentialsHandler)
		subRouter.Get("/annual", statsHandler)
		subRouter.Get("/last-activity", lastActivityHandler)
	})
}

func NewStravaWidget() *StravaWidget {
	return &StravaWidget{}
}
