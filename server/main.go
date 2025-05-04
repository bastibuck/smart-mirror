package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// Define a struct for the sports data
type SportStats struct {
	Count      int `json:"count"`
	MovingTime int `json:"moving_time"`
	Distance   int `json:"distance"`
}

// Constructor function for SportStats
func NewSportStats(count, movingTime, distance int) SportStats {
	return SportStats{
		Count:      count,
		MovingTime: movingTime,
		Distance:   distance,
	}
}

// Define a struct for the overall data
type StravaStats struct {
	Running SportStats `json:"running"`
	Cycling SportStats `json:"cycling"`
	Kiting  SportStats `json:"kiting"`
}

func main() {
	fmt.Println("Starting the application...")

	router := setupRouter()

	// Start the server
	http.ListenAndServe(":8080", router)
}

func setupRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
		MaxAge:         300,
	}))

	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World!"))
	})

	router.Get("/strava-stats", func(res http.ResponseWriter, req *http.Request) {
		data := StravaStats{
			Running: NewSportStats(10, 200, 186),
			Cycling: NewSportStats(20, 300, 305),
			Kiting:  NewSportStats(30, 400, 237),
		}

		// Set the response header to JSON
		res.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(res).Encode(data); err != nil {
			http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
		}
	})

	return router
}
