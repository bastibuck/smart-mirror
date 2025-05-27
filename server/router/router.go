package router

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"smartmirror.server/env"
	"smartmirror.server/router/routes"
)

func SetupRouter() *chi.Mux {
	router := chi.NewRouter()

	// Middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{os.Getenv(env.EnvCorsAllowedOrigin)},
		AllowedMethods: []string{"GET", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
		MaxAge:         300,
	}))

	// Route Groups
	routes.RegisterHomeRoutes(router)
	routes.RegisterStravaRoutes(router)
	routes.RegisterVersionHashRoutes(router)

	return router
}
