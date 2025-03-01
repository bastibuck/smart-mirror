package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Starting the application...")

	router := chi.NewRouter()

	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World!"))
	})

	// Start the server
	http.ListenAndServe(":8080", router)
}
