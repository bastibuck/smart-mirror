package strava

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var GLOBAL_StravaAthleteId int
var GLOBAL_ExpiresAt int
var GLOBAL_StravaAccessToken string
var GLOBAL_StravaRefreshToken string

type stravaAPIResponse struct {
	YtdRideTotals struct {
		Count      int     `json:"count"`
		Distance   int     `json:"distance"`
		MovingTime float32 `json:"moving_time"`
	} `json:"ytd_ride_totals"`
	YtdRunTotals struct {
		Count      int     `json:"count"`
		Distance   int     `json:"distance"`
		MovingTime float32 `json:"moving_time"`
	} `json:"ytd_run_totals"`
}

func FetchStravaData() (stravaStats, error) {
	if cachedData, found := getCachedStravaStats(); found {
		return cachedData, nil
	}

	// Check if the global Strava athlete ID and access token are set
	if GLOBAL_StravaAthleteId == 0 || GLOBAL_StravaAccessToken == "" || GLOBAL_StravaRefreshToken == "" {
		return stravaStats{}, fmt.Errorf("401") // TODO? make this return directly instead of passing outside as string?
	}

	err := refreshStravaAccessToken()

	if err != nil {
		return stravaStats{}, fmt.Errorf("401") // TODO? make this return directly instead of passing outside as string?
	}

	stravaAPIURL := fmt.Sprintf("https://www.strava.com/api/v3/athletes/%d/stats", GLOBAL_StravaAthleteId)

	req, err := http.NewRequest("GET", stravaAPIURL, nil)
	if err != nil {
		return stravaStats{}, err
	}

	req.Header.Set("Authorization", "Bearer "+GLOBAL_StravaAccessToken)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return stravaStats{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			return stravaStats{}, fmt.Errorf("%d", resp.StatusCode) // TODO? make this return directly instead of passing outside as string?
		}

		return stravaStats{}, fmt.Errorf("Strava API returned status: %s", resp.Status)
	}

	var stravaAPIResponse stravaAPIResponse

	if err := json.NewDecoder(resp.Body).Decode(&stravaAPIResponse); err != nil {
		return stravaStats{}, err
	}

	stats := stravaStats{
		Cycling: sportStats{
			Count:       stravaAPIResponse.YtdRideTotals.Count,
			MovingTimeS: int(stravaAPIResponse.YtdRideTotals.MovingTime),
			DistanceM:   stravaAPIResponse.YtdRideTotals.Distance,
		},
		Running: sportStats{
			Count:       stravaAPIResponse.YtdRunTotals.Count,
			MovingTimeS: int(stravaAPIResponse.YtdRunTotals.MovingTime),
			DistanceM:   stravaAPIResponse.YtdRunTotals.Distance,
		},
	}

	setCachedStravaStats(stats)

	return stats, nil
}

type stravaRefreshTokenApiResponse struct {
	// TokenType    string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresAt   int    `json:"expires_at"`
	// ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func refreshStravaAccessToken() error {
	// Check if the access token is still valid
	if (GLOBAL_ExpiresAt - 1*int(time.Hour.Seconds())) > int(time.Now().Unix()) {
		return nil
	}

	url := "https://www.strava.com/oauth/token" +
		"?client_id=" + getStravaClientId() +
		"&client_secret=" + getStravaClientSecret() +
		"&refresh_token=" + GLOBAL_StravaRefreshToken +
		"&grant_type=refresh_token"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Strava API returned status: %s", resp.Status)
	}

	var response stravaRefreshTokenApiResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	GLOBAL_StravaAccessToken = response.AccessToken
	GLOBAL_StravaRefreshToken = response.RefreshToken
	GLOBAL_ExpiresAt = response.ExpiresAt

	return nil
}

type stravaExchangeTokenAPIResponse struct {
	ExpiresAt    int    `json:"expires_at"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Athlete      struct {
		Id int `json:"id"`
	}
}

func ExchangeCodeForToken(code string) error {
	url := "https://www.strava.com/oauth/token" +
		"?client_id=" + getStravaClientId() +
		"&client_secret=" + getStravaClientSecret() +
		"&code=" + code +
		"&grant_type=authorization_code"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("Failed to create request: %v", err)
	}

	// call strava api to exchange token
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to exchange token: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Strava API returned status: %s", resp.Status)
	}

	var response stravaExchangeTokenAPIResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("Failed decode response: %v", err)
	}

	GLOBAL_StravaAccessToken = response.AccessToken
	GLOBAL_StravaRefreshToken = response.RefreshToken
	GLOBAL_StravaAthleteId = response.Athlete.Id
	GLOBAL_ExpiresAt = response.ExpiresAt

	return nil
}
