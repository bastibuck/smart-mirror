package strava

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func statsHandler(res http.ResponseWriter, req *http.Request) {
	stats, err := fetchStravaData()

	if err != nil {
		if err.Error() == "401" {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
			return
		}

		logger.Info("Failed to fetch data from Strava: %v", err)
		http.Error(res, fmt.Sprintf("Failed to fetch data from Strava: %v", err), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(res).Encode(stats); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func lastActivityHandler(res http.ResponseWriter, req *http.Request) {
	lastActivity, err := fetchLastActivity()

	if err != nil {
		if err.Error() == "401" {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
			return
		}

		logger.Info("Failed to fetch last activity from Strava: %v", err)
		http.Error(res, fmt.Sprintf("Failed to fetch last activity from Strava: %v", err), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(res).Encode(lastActivity); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func exchangeTokenHandler(res http.ResponseWriter, req *http.Request) {
	err := exchangeCodeForToken(req.URL.Query().Get("code"))

	if err != nil {
		http.Redirect(res, req, getStravaLoginFailureUrl(), http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(res, req, getStravaLoginSuccessUrl(), http.StatusTemporaryRedirect)
}

func credentialsHandler(res http.ResponseWriter, req *http.Request) {
	creds, err := getStravaCredentials()

	if err != nil {
		http.Error(res, fmt.Sprintf("Failed to get Strava credentials: %v", err), http.StatusForbidden)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(res).Encode(creds); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
