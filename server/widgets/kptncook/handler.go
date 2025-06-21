package kptncook

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func dailyRecipesHandler(w http.ResponseWriter, r *http.Request) {
	response, err := getDailyRecipes()

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get daily recipes: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
