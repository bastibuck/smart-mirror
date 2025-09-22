package windspeed

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func windspeedHandler(res http.ResponseWriter, req *http.Request) {
	windspeedData, err := getCurrentWind()

	if err != nil {
		http.Error(res, fmt.Sprintf("Failed to get current wind: %v", err), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(res).Encode(windspeedData); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
