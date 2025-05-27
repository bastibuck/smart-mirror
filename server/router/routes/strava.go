package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"smartmirror.server/strava"
)

func RegisterStravaRoutes(router *chi.Mux) {
	router.Get("/strava/stats", stravaStatsHandler)
	router.Get("/strava/exchange-token", stravaExchangeTokenHandler)
}

func stravaStatsHandler(res http.ResponseWriter, req *http.Request) {
	stravaResponse, err := strava.FetchStravaData()

	if err != nil {
		if err.Error() == "401" {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
			return
		}

		http.Error(res, fmt.Sprintf("Failed to fetch data from Strava: %v", err), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(res).Encode(stravaResponse); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func stravaExchangeTokenHandler(res http.ResponseWriter, req *http.Request) {
	err := strava.ExchangeCodeForToken(req.URL.Query().Get("code"))

	if err != nil {
		http.Redirect(res, req, "http://TODO", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(res, req, strava.GetStravaLoginSuccessUrl(), http.StatusTemporaryRedirect)
}
