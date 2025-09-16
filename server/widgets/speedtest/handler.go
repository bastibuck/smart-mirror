package speedtest

import (
	"encoding/json"
	"net/http"
	"time"
)

func speedtestHandler(res http.ResponseWriter, req *http.Request) {
	speedtestResult := getSpeedTestResults(-4 * time.Hour)

	res.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(res).Encode(speedtestResult); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
