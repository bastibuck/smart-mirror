package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func SetupRouter() *chi.Mux {
	setupEnv()

	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{getCorsAllowedOrigin()},
		AllowedMethods: []string{"GET", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
		MaxAge:         300,
	}))

	// Root
	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Smart mirror server is running!"))
	})

	return router
}
