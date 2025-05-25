package routes

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"smartmirror.server/config"
	routes "smartmirror.server/routes/strava"
)

func SetupRouter() *chi.Mux {
	corsAllowedOrigin := os.Getenv(config.EnvCorsAllowedOrigin)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{corsAllowedOrigin},
		AllowedMethods: []string{"GET", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
		MaxAge:         300,
	}))

	router.Get("/", HomeHandler)
	router.Get("/strava-stats", routes.StravaStatsHandler)
	router.Get("/strava-stats/exchange-token", routes.StravaExchangeTokenHandler)
	router.Get("/version-hash", VersionHashHandler)

	return router
}
