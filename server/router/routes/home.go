package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterHomeRoutes(router *chi.Mux) {
	router.Get("/", homeHandler)
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Smart mirror server is running!"))
}
