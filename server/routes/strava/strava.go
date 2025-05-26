package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
	"smartmirror.server/config"
)

var GLOBAL_StravaAthleteId int
var GLOBAL_StravaAccessToken string
var GLOBAL_StravaRefreshToken string

type SportStats struct {
	Count       int `json:"count"`
	MovingTimeS int `json:"moving_time_s"`
	DistanceM   int `json:"distance_m"`
}

type StravaStats struct {
	Running SportStats `json:"running"`
	Cycling SportStats `json:"cycling"`
}

func NewSportStats(count, movingTimeS, distanceM int) SportStats {
	return SportStats{
		Count:       count,
		MovingTimeS: movingTimeS,
		DistanceM:   distanceM,
	}
}

type StravaAPIResponse struct {
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

var stravaCache = cache.New(30*time.Minute, 45*time.Minute)

func StravaStatsHandler(res http.ResponseWriter, req *http.Request) {
	stravaResponse, err := fetchStravaData()

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

func fetchStravaData() (StravaStats, error) {
	const cacheKey = "strava-stats"

	if cachedData, found := stravaCache.Get(cacheKey); found {
		return cachedData.(StravaStats), nil
	}

	// Check if the global Strava athlete ID and access token are set
	if GLOBAL_StravaAthleteId == 0 || GLOBAL_StravaAccessToken == "" {
		return StravaStats{}, fmt.Errorf("401") // TODO? make this return directly instead of passing outside as string?
	}

	err := refreshStravaAccessToken()

	if err != nil {
		return StravaStats{}, fmt.Errorf("failed to refresh Strava access token: %v", err)
	}

	stravaAPIURL := fmt.Sprintf("https://www.strava.com/api/v3/athletes/%d/stats", GLOBAL_StravaAthleteId)

	req, err := http.NewRequest("GET", stravaAPIURL, nil)
	if err != nil {
		return StravaStats{}, err
	}

	req.Header.Set("Authorization", "Bearer "+GLOBAL_StravaAccessToken)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return StravaStats{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			return StravaStats{}, fmt.Errorf("%d", resp.StatusCode) // TODO? make this return directly instead of passing outside as string?
		}

		return StravaStats{}, fmt.Errorf("Strava API returned status: %s", resp.Status)
	}

	var stravaAPIResponse StravaAPIResponse

	if err := json.NewDecoder(resp.Body).Decode(&stravaAPIResponse); err != nil {
		return StravaStats{}, err
	}

	stats := StravaStats{
		Cycling: NewSportStats(
			stravaAPIResponse.YtdRideTotals.Count,
			int(stravaAPIResponse.YtdRideTotals.MovingTime),
			stravaAPIResponse.YtdRideTotals.Distance,
		),
		Running: NewSportStats(
			stravaAPIResponse.YtdRunTotals.Count,
			int(stravaAPIResponse.YtdRunTotals.MovingTime),
			stravaAPIResponse.YtdRunTotals.Distance,
		),
	}

	stravaCache.Set(cacheKey, stats, cache.DefaultExpiration)

	return stats, nil
}

type StravaRefreshTokenApiResponse struct {
	// TokenType    string `json:"token_type"`
	AccessToken string `json:"access_token"`
	// ExpiresAt    int    `json:"expires_at"`
	// ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func refreshStravaAccessToken() error {
	url := "https://www.strava.com/oauth/token" +
		"?client_id=" + os.Getenv(config.EnvStravaClientId) +
		"&client_secret=" + os.Getenv(config.EnvStravaClientSecret) +
		"&grant_type=refresh_token" +
		"&refresh_token=" + GLOBAL_StravaRefreshToken

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

	var response StravaRefreshTokenApiResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	GLOBAL_StravaAccessToken = response.AccessToken
	GLOBAL_StravaRefreshToken = response.RefreshToken

	return nil
}
