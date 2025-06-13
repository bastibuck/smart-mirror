package garmin

import (
	"encoding/json"
	"net/http"
)

type stepsTodayHandlerResponse struct {
	Steps int `json:"steps"`
}

func stepsTodayHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	steps, err := getSevenDaySteps()

	if err != nil {
		http.Error(res, "Failed to get steps", http.StatusInternalServerError)
		return
	}

	response := stepsTodayHandlerResponse{
		Steps: steps.Days[len(steps.Days)-1].Steps,
	}

	if err := json.NewEncoder(res).Encode(response); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

type stepsThisWeekHandlerResponse struct {
	Total   int `json:"total"`
	Average int `json:"average"`
	Days    []struct {
		Date  string `json:"date"`
		Steps string `json:"steps"`
	} `json:"days"`
}

func stepsThisWeekHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	steps, err := getSevenDaySteps()

	if err != nil {
		http.Error(res, "Failed to get steps", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(res).Encode(steps); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
