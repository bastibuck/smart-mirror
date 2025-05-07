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

type SportStats struct {
	Count       int `json:"count"`
	MovingTimeS int `json:"moving_time_s"`
	Distance    int `json:"distance"`
}

type StravaStats struct {
	Running SportStats `json:"running"`
	Cycling SportStats `json:"cycling"`
}

func NewSportStats(count, movingTimeS, distance int) SportStats {
	return SportStats{
		Count:       count,
		MovingTimeS: movingTimeS,
		Distance:    distance,
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

var stravaCache = cache.New(30*time.Minute, 45*time.Hour)

func StravaStatsHandler(res http.ResponseWriter, req *http.Request) {
	stravaResponse, err := fetchStravaData()
	if err != nil {
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

	athleteID := os.Getenv(config.EnvStravaAthleteID)
	accessToken := os.Getenv(config.EnvStravaAccessToken)

	stravaAPIURL := fmt.Sprintf("https://www.strava.com/api/v3/athletes/%s/stats", athleteID)

	req, err := http.NewRequest("GET", stravaAPIURL, nil)
	if err != nil {
		return StravaStats{}, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return StravaStats{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
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
