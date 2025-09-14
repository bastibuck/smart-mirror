package speedtest

import (
	"encoding/json"
	"net/http"
)

func speedtestHandler(res http.ResponseWriter, req *http.Request) {

	speedtestResult, err := runSpeedtest()

	if err != nil {
		http.Error(res, "Failed to run speedtest", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(res).Encode(speedtestResult); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
