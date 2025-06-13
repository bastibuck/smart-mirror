package garmin

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterVersionHashRoutes(router *chi.Mux) {
	router.HandleFunc("/steps/today", stepsTodayHandler)
}

type stepsTodayhHandlerResponse struct {
	Steps int `json:"steps"`
}

func stepsTodayHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	steps, err := getStepsToday()

	if err != nil {
		http.Error(res, "Failed to get steps", http.StatusInternalServerError)
		return
	}

	response := stepsTodayhHandlerResponse{
		Steps: steps,
	}

	if err := json.NewEncoder(res).Encode(response); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
