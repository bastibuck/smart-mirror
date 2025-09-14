package speedtest

import (
	"encoding/json"
	"net/http"
)

type speedtestHandlerResponse struct {
	Download float64 `json:"download"`
	Upload   float64 `json:"upload"`
	Ping     float64 `json:"ping"`
}

func speedtestHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	response := speedtestHandlerResponse{
		Download: 0,
		Upload:   0,
		Ping:     0,
	}

	if err := json.NewEncoder(res).Encode(response); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
