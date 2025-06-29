package strava

import (
	"github.com/go-chi/chi/v5"
	"smartmirror.server/env"
	"smartmirror.server/widgets"
)

type StravaWidget struct{}

var _ widgets.Widget = (*StravaWidget)(nil)

func (v *StravaWidget) SetupEnv() {
	setDefaultEnv()
	env.ValidateEnvKeys(getEnvKeys())
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
