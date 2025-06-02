package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"smartmirror.server/router/routes"
	"smartmirror.server/shared"
)

func SetupRouter() *chi.Mux {
	router := chi.NewRouter()

	// Middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{shared.GetCorsAllowedOrigin()},
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
