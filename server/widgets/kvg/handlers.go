package kvg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func nextDeparturesHandler(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

	if err != nil || limit <= 0 {
		limit = 5 // Default limit if not specified or invalid
	}

	nextDepartures, err := fetchNextDepartures(limit) // Default limit to 5 if not specified)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch next departures from KVG: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(nextDepartures); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
