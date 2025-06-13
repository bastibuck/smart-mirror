package strava

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/twpayne/go-polyline"
	"smartmirror.server/widgets/shared"
)

var GLOBAL_ExpiresAt int
var GLOBAL_StravaAccessToken string
var GLOBAL_StravaRefreshToken string

type athleteActivityResponseModel struct {
	Name       string  `json:"name"`
	SportType  string  `json:"sport_type"`
	Distance   float32 `json:"distance"`    // in meters
	MovingTime float32 `json:"moving_time"` // in seconds
	Map        struct {
		SummaryPolyline string `json:"summary_polyline"`
	} `json:"map"`
	StartDate string `json:"start_date_local"`
}

func fetchStravaData() (annualStatsModel, error) {
	if cachedData, found := stravaCache.GetAnnualStats(); found {
		return cachedData, nil
	}

	// Check if the global Strava athlete ID and access token are set
	if GLOBAL_StravaAccessToken == "" || GLOBAL_StravaRefreshToken == "" {
		return annualStatsModel{}, fmt.Errorf("401") // TODO? make this return directly instead of passing outside as string?
	}

	err := refreshStravaAccessToken()

	if err != nil {
		return annualStatsModel{}, fmt.Errorf("401") // TODO? make this return directly instead of passing outside as string?
	}

	var BEGINNING_OF_YEAR int64 = time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.UTC).Unix()

	var responseBucket []athleteActivityResponseModel
	page := 0
	maxRequests := 20 // maximum number of requests to Strava API
	for {
		page++

		stravaAPIURL := fmt.Sprintf("https://www.strava.com/api/v3/athlete/activities?after=%d&page=%d&per_page=200", BEGINNING_OF_YEAR, page)

		req, err := http.NewRequest("GET", stravaAPIURL, nil)
		if err != nil {
			return annualStatsModel{}, err
		}

		req.Header.Set("Authorization", "Bearer "+GLOBAL_StravaAccessToken)

		// Make the HTTP request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return annualStatsModel{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			if resp.StatusCode == http.StatusUnauthorized {
				return annualStatsModel{}, fmt.Errorf("%d", resp.StatusCode) // TODO? make this return directly instead of passing outside as string?
			}

			return annualStatsModel{}, fmt.Errorf("Strava API returned status: %s", resp.Status)
		}

		var response []athleteActivityResponseModel

		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return annualStatsModel{}, err
		}

		responseBucket = append(responseBucket, response...)

		if len(response) == 0 || page >= maxRequests {
			break
		}
	}

	Running := sportStatsModel{}
	Hiking := sportStatsModel{}
	Cycling := sportStatsModel{}
	Kiting := sportStatsModel{}

	statsMap := map[string]*sportStatsModel{
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

	stats := annualStatsModel{
		Running: Running,
		Hiking:  Hiking,
		Cycling: Cycling,
		Kiting:  Kiting,
	}

	stravaCache.SetAnnualStats(stats)

	return stats, nil
}

func fetchLastActivity() (lastActivityModel, error) {
	if cachedData, found := stravaCache.GetLastActivity(); found {
		return cachedData, nil
	}

	if GLOBAL_StravaAccessToken == "" || GLOBAL_StravaRefreshToken == "" {
		return lastActivityModel{}, fmt.Errorf("401") // TODO? make this return directly instead of passing outside as string?
	}

	err := refreshStravaAccessToken()

	if err != nil {
		return lastActivityModel{}, fmt.Errorf("401") // TODO? make this return directly instead of passing outside as string?
	}

	stravaAPIURL := "https://www.strava.com/api/v3/athlete/activities?per_page=1"

	req, err := http.NewRequest("GET", stravaAPIURL, nil)
	if err != nil {
		return lastActivityModel{}, err
	}

	req.Header.Set("Authorization", "Bearer "+GLOBAL_StravaAccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return lastActivityModel{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			return lastActivityModel{}, fmt.Errorf("%d", resp.StatusCode) // TODO? make this return directly instead of passing outside as string?
		}

		return lastActivityModel{}, fmt.Errorf("Strava API returned status: %s", resp.Status)
	}

	var response []athleteActivityResponseModel

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return lastActivityModel{}, err
	}

	if len(response) == 0 {
		return lastActivityModel{}, fmt.Errorf("No activities found")
	}

	activity := response[0]

	typeMap := map[string]string{
		"Run":        "Run",
		"VirtualRun": "Run",
		"TrailRun":   "Run",

		"Ride":              "Ride",
		"GravelRide":        "Ride",
		"VirtualRide":       "Ride",
		"MountainBikeRide":  "Ride",
		"EBikeRide":         "Ride",
		"EMountainBikeRide": "Ride",

		"Kitesurf": "Kite",

		"Hike": "Hike",
	}

	buf := []byte(activity.Map.SummaryPolyline)
	rawCoords, _, err := polyline.DecodeCoords(buf)

	if err != nil {
		return lastActivityModel{}, fmt.Errorf("Failed to decode polyline: %v", err)
	}

	normalizedCoords := make([][]float64, len(rawCoords))
	for i, coord := range rawCoords {
		normalizedCoords[i] = []float64{coord[1], coord[0]}
	}

	lastActivityData := lastActivityModel{
		Name:        activity.Name,
		Date:        activity.StartDate,
		Coordinates: normalizedCoords,
		Type:        typeMap[activity.SportType],
		DistanceM:   int(activity.Distance),
		MovingTimeS: int(activity.MovingTime),
	}

	stravaCache.SetLastActivity(lastActivityData)

	return lastActivityData, nil
}

type refreshTokenResponseModel struct {
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

	var response refreshTokenResponseModel

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	GLOBAL_StravaAccessToken = response.AccessToken
	GLOBAL_StravaRefreshToken = response.RefreshToken
	GLOBAL_ExpiresAt = response.ExpiresAt

	return nil
}

type exchangeTokenResponseModel struct {
	ExpiresAt    int    `json:"expires_at"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Athlete      struct {
		Id int `json:"id"`
	}
}

func exchangeCodeForToken(code string) error {
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

	var response exchangeTokenResponseModel

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("Failed decode response: %v", err)
	}

	GLOBAL_StravaAccessToken = response.AccessToken
	GLOBAL_StravaRefreshToken = response.RefreshToken
	GLOBAL_ExpiresAt = response.ExpiresAt

	return nil
}

func getStravaCredentials() (credentialsModel, error) {
	if shared.GetAppMode() != "development" {
		return credentialsModel{}, fmt.Errorf("Strava credentials are only available in development mode")
	}

	if GLOBAL_StravaAccessToken == "" || GLOBAL_StravaRefreshToken == "" || GLOBAL_ExpiresAt == 0 {
		return credentialsModel{}, fmt.Errorf("Strava credentials are not set")
	}

	return credentialsModel{
		AccessToken:  GLOBAL_StravaAccessToken,
		RefreshToken: GLOBAL_StravaRefreshToken,
		ExpiresAt:    GLOBAL_ExpiresAt,
	}, nil
}
