package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

// Define a struct for the sports data
type SportStats struct {
	Count       int `json:"count"`
	MovingTimeS int `json:"moving_time_s"`
	Distance    int `json:"distance"`
}

// Constructor function for SportStats
func NewSportStats(count, movingTimeS, distance int) SportStats {
	return SportStats{
		Count:       count,
		MovingTimeS: movingTimeS,
		Distance:    distance,
	}
}

// Define a struct for the overall data
type StravaStats struct {
	Running SportStats `json:"running"`
	Cycling SportStats `json:"cycling"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	// Validate required environment variables
	requiredKeys := []string{"STRAVA_ACCESS_TOKEN", "STRAVA_ATHLETE_ID"}
	for _, key := range requiredKeys {
		if os.Getenv(key) == "" {
			fmt.Printf("Missing required environment variable: %s\n", key)
			os.Exit(1)
		}
	}

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
		// TODO? cache?

		stravaResponse, err := fetchStravaData()
		if err != nil {
			http.Error(res, fmt.Sprintf("Failed to fetch data from Strava: %v", err), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(res).Encode(stravaResponse); err != nil {
			http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
		}
	})

	return router
}

func fetchStravaData() (StravaStats, error) {
	athleteID := os.Getenv("STRAVA_ATHLETE_ID")
	accessToken := os.Getenv("STRAVA_ACCESS_TOKEN")

	stravaAPIURL := fmt.Sprintf("https://www.strava.com/api/v3/athletes/%s/stats", athleteID)

	req, err := http.NewRequest("GET", stravaAPIURL, nil)
	if err != nil {
		return StravaStats{}, err
	}

	// Add the Authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return StravaStats{}, err
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return StravaStats{}, fmt.Errorf("Strava API returned status: %s", resp.Status)
	}

	// Parse the response body
	var stravaAPIResponse struct {
		YtdRideTotals struct {
			Count      int     `json:"count"`
			Distance   int     `json:"distance"`
			MovingTime float32 `json:"moving_time"`
		} `json:"ytd_ride_totals"`
		YtdRunTotals struct {
			Count      int     `json:"count"`
			Distance   int     `json:"distance"`
			MovingTime float32 `json:"moving_time"`
		} `json:"ytd_run_totals"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&stravaAPIResponse); err != nil {
		return StravaStats{}, err
	}

	// Map the Strava API response to your StravaStats struct
	return StravaStats{
		Cycling: NewSportStats(
			stravaAPIResponse.YtdRideTotals.Count,
			int(stravaAPIResponse.YtdRideTotals.MovingTime),
			stravaAPIResponse.YtdRideTotals.Distance,
		),
		Running: NewSportStats(
			stravaAPIResponse.YtdRunTotals.Count,
			int(stravaAPIResponse.YtdRunTotals.MovingTime),
			stravaAPIResponse.YtdRunTotals.Distance,
		),
	}, nil
}
