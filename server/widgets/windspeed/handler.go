package windspeed

import (
	"encoding/json"
	"net/http"
)

type WindspeedData struct {
	SpeedKn      float64 `json:"speed_kn"`
	DirectionDeg int16   `json:"direction_deg"`
}

func windspeedHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(res).Encode(WindspeedData{}); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
