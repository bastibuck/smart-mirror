package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Starting the application...")

	router := setupRouter()

	// Start the server
	http.ListenAndServe(":8080", router)
}

func setupRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World!"))
	})

	return router
}
