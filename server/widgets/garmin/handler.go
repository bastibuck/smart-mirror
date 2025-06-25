package garmin

import (
	"encoding/json"
	"net/http"

	"github.com/bastibuck/go-garmin"
)

type stepsThisWeekHandlerResponse struct {
	Total   int `json:"total"`
	Average int `json:"average"`
	Days    []struct {
		Date  string `json:"date"`
		Steps string `json:"steps"`
	} `json:"days"`
}

func stepsThisWeekHandler(res http.ResponseWriter, apiClient *garmin.API) {
	res.Header().Set("Content-Type", "application/json")

	steps, err := getSevenDaySteps(apiClient)

	if err != nil {
		http.Error(res, "Failed to get steps", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(res).Encode(steps); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
