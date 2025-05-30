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

	// Hikes
	hikeActivities := 0
	hikeDistanceM := float32(0)
	hikeMovingTimeS := float32(0)

	// Running
	runningActivities := 0
	runningDistanceM := float32(0)
	runningMovingTimeS := float32(0)

	// Kiting
	kiteActivities := 0
	kiteDistanceM := float32(0)
	kiteMovingTimeS := float32(0)

	// Cycling
	cyclingActivities := 0
	cyclingDistanceM := float32(0)
	cyclingMovingTimeS := float32(0)

	for _, activity := range responseBucket {
		if activity.SportType == "Run" ||
			activity.SportType == "VirtualRun" ||
			activity.SportType == "TrailRun" {
			runningActivities++
			runningDistanceM += activity.Distance
			runningMovingTimeS += activity.MovingTime
		}

		if activity.SportType == "Ride" ||
			activity.SportType == "GravelRide" ||
			activity.SportType == "VirtualRide" ||
			activity.SportType == "MountainBikeRide" ||
			activity.SportType == "EBikeRide" ||
			activity.SportType == "EMountainBikeRide" {
			cyclingActivities++
			cyclingDistanceM += activity.Distance
			cyclingMovingTimeS += activity.MovingTime
		}

		if activity.SportType == "Hike" {
			hikeActivities++
			hikeDistanceM += activity.Distance
			hikeMovingTimeS += activity.MovingTime
		}

		if activity.SportType == "Kitesurf" {
			kiteActivities++
			kiteDistanceM += activity.Distance
			kiteMovingTimeS += activity.MovingTime
		}
	}

	stats := stravaStats{
		Running: sportStats{
			Count:       runningActivities,
			DistanceM:   int(runningDistanceM),
			MovingTimeS: int(runningMovingTimeS),
		},

		Cycling: sportStats{
			Count:       cyclingActivities,
			DistanceM:   int(cyclingDistanceM),
			MovingTimeS: int(cyclingMovingTimeS),
		},

		Kiting: sportStats{
			Count:       kiteActivities,
			DistanceM:   int(kiteDistanceM),
			MovingTimeS: int(kiteMovingTimeS),
		},

		Hiking: sportStats{
			Count:       hikeActivities,
			DistanceM:   int(hikeDistanceM),
			MovingTimeS: int(hikeMovingTimeS),
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
	GLOBAL_ExpiresAt = response.ExpiresAt

	return nil
}
