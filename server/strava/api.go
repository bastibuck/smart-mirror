package strava

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var GLOBAL_ExpiresAt int
var GLOBAL_StravaAccessToken string
var GLOBAL_StravaRefreshToken string

type stravaAPIResponse struct {
	Name       string  `json:"name"`
	SportType  string  `json:"sport_type"`
	Distance   float32 `json:"distance"`    // in meters
	MovingTime float32 `json:"moving_time"` // in seconds
}

func FetchStravaData() (stravaStats, error) {
	if cachedData, found := getCachedStravaStats(); found {
		return cachedData, nil
	}

	// Check if the global Strava athlete ID and access token are set
	if GLOBAL_StravaAccessToken == "" || GLOBAL_StravaRefreshToken == "" {
		return stravaStats{}, fmt.Errorf("401") // TODO? make this return directly instead of passing outside as string?
	}

	err := refreshStravaAccessToken()

	if err != nil {
		return stravaStats{}, fmt.Errorf("401") // TODO? make this return directly instead of passing outside as string?
	}

	var BEGINNING_OF_YEAR int64 = time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.UTC).Unix()

	var responseBucket []stravaAPIResponse
	page := 0
	maxRequests := 20 // maximum number of requests to Strava API
	for {
		page++

		stravaAPIURL := fmt.Sprintf("https://www.strava.com/api/v3/athlete/activities?after=%d&page=%d&per_page=200", BEGINNING_OF_YEAR, page)

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

		var response []stravaAPIResponse

		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return stravaStats{}, err
		}

		responseBucket = append(responseBucket, response...)

		if len(response) == 0 || page >= maxRequests {
			break
		}
	}

	Running := sportStats{}
	Hiking := sportStats{}
	Cycling := sportStats{}
	Kiting := sportStats{}

	statsMap := map[string]*sportStats{
		"Run":        &Running,
		"VirtualRun": &Running,
		"TrailRun":   &Running,

		"Ride":              &Cycling,
		"GravelRide":        &Cycling,
		"VirtualRide":       &Cycling,
		"MountainBikeRide":  &Cycling,
		"EBikeRide":         &Cycling,
		"EMountainBikeRide": &Cycling,

		"Kitesurf": &Kiting,

		"Hike": &Hiking,
	}

	for _, activity := range responseBucket {
		if stat, ok := statsMap[activity.SportType]; ok {
			stat.Count++
			stat.DistanceM += int(activity.Distance)
			stat.MovingTimeS += int(activity.MovingTime)
		}
	}

	stats := stravaStats{
		Running: Running,
		Hiking:  Hiking,
		Cycling: Cycling,
		Kiting:  Kiting,
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
	GLOBAL_ExpiresAt = response.ExpiresAt

	return nil
}
