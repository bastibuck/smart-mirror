package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"smartmirror.server/config"
)

type StravaExchangeTokenAPIResponse struct {
	ExpiresAt    int    `json:"expires_at"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Athlete      struct {
		Id int `json:"id"`
	}
}

func StravaExchangeTokenHandler(res http.ResponseWriter, req *http.Request) {
	url := "https://www.strava.com/oauth/token?client_id=" + os.Getenv(config.EnvStravaClientId) +
		"&client_secret=" + os.Getenv(config.EnvStravaClientSecret) +
		"&code=" + req.URL.Query().Get("code") +
		"&grant_type=authorization_code"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		http.Error(res, fmt.Sprintf("Failed to create request: %v", err), http.StatusInternalServerError)
		return
	}

	// call strava api to exchange token
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(res, fmt.Sprintf("Failed to exchange token: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(res, fmt.Sprintf("Strava API returned status: %s", resp.Status), http.StatusInternalServerError)
		return
	}

	var response StravaExchangeTokenAPIResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		http.Error(res, fmt.Sprintf("Failed decode response: %v", err), http.StatusInternalServerError)
		return
	}

	GLOBAL_StravaAccessToken = response.AccessToken
	GLOBAL_StravaRefreshToken = response.RefreshToken
	GLOBAL_StravaAthleteId = response.Athlete.Id

	// redirect user to success
	http.Redirect(res, req, os.Getenv(config.EnvFrontendUrl)+"/strava/token-success", http.StatusTemporaryRedirect)
}
