package strava

import (
	"fmt"
	"time"

	"github.com/twpayne/go-polyline"
	"smartmirror.server/utils"
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
	if cachedData, found := stravaCache.getAnnualStats(); found {
		logger.Info("Using cached data for annual data")
		return cachedData, nil
	}

	// Check if the global Strava athlete ID and access token are set
	if GLOBAL_StravaAccessToken == "" || GLOBAL_StravaRefreshToken == "" {
		return annualStatsModel{}, fmt.Errorf("401") // TODO? make this return directly instead of passing outside as string?
	}

	err := refreshStravaAccessToken()

	if err != nil {
		return annualStatsModel{}, fmt.Errorf("fetchStravaData: Failed refreshing access token: %v", err)
	}

	var BEGINNING_OF_YEAR int64 = time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.UTC).Unix()

	var responseBucket []athleteActivityResponseModel
	page := 0
	maxRequests := 20 // maximum number of requests to Strava API
	for {
		page++

		stravaAnnualActivitiesUrl := fmt.Sprintf("https://www.strava.com/api/v3/athlete/activities?after=%d&page=%d&per_page=200", BEGINNING_OF_YEAR, page)

		var response []athleteActivityResponseModel

		err = utils.RelaxedHttpRequest(utils.RelaxedHttpRequestOptions{
			URL:      stravaAnnualActivitiesUrl,
			Response: &response,
			Headers: map[string]string{
				"Authorization": "Bearer " + GLOBAL_StravaAccessToken,
			},
		})

		if err != nil {
			if err.Error() == "401" {
				return annualStatsModel{}, fmt.Errorf("401") // TODO? make this return directly instead of passing outside as string?
			}

			return annualStatsModel{}, fmt.Errorf("Failed to fetch annual activities from strava: %v", err)
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

	stravaCache.setAnnualStats(stats)

	return stats, nil
}

func fetchLastActivity() (lastActivityModel, error) {
	if cachedData, found := stravaCache.getLastActivity(); found {
		logger.Info("Using cached data for last activity: %v", cachedData.Name)
		return cachedData, nil
	}

	if GLOBAL_StravaAccessToken == "" || GLOBAL_StravaRefreshToken == "" {
		return lastActivityModel{}, fmt.Errorf("401") // TODO? make this return directly instead of passing outside as string?
	}

	err := refreshStravaAccessToken()

	if err != nil {
		return lastActivityModel{}, fmt.Errorf("fetchLastActivity: Failed to fresh access token: %v", err)
	}

	var response []athleteActivityResponseModel

	err = utils.RelaxedHttpRequest(utils.RelaxedHttpRequestOptions{
		URL:      "https://www.strava.com/api/v3/athlete/activities?per_page=1",
		Response: &response,
		Headers: map[string]string{
			"Authorization": "Bearer " + GLOBAL_StravaAccessToken,
		},
	})

	if err != nil {
		if err.Error() == "401" {
			return lastActivityModel{}, fmt.Errorf("401") // TODO? make this return directly instead of passing outside as string?
		}

		return lastActivityModel{}, fmt.Errorf("Failed to fetch last activity from strava: %v", err)
	}

	if len(response) == 0 {
		return lastActivityModel{}, fmt.Errorf("No activities found")
	}

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

	// find first supported activity
	var activity *athleteActivityResponseModel = nil
	for i := range response {
		if _, ok := typeMap[response[i].SportType]; ok {
			activity = &response[i]
			break
		}
	}

	if activity == nil {
		return lastActivityModel{}, fmt.Errorf("No supported activity found")
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

	stravaCache.setLastActivity(lastActivityData)

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

	logger.Info("Access token expired, getting a new one.")

	var response refreshTokenResponseModel

	url := "https://www.strava.com/oauth/token" +
		"?client_id=" + getStravaClientId() +
		"&client_secret=" + getStravaClientSecret() +
		"&refresh_token=" + GLOBAL_StravaRefreshToken +
		"&grant_type=refresh_token"

	err := utils.RelaxedHttpRequest(utils.RelaxedHttpRequestOptions{
		URL:      url,
		Method:   "POST",
		Response: &response,
	})

	if err != nil {
		if err.Error() == "401" {
			return fmt.Errorf("401") // TODO? make this return directly instead of passing outside as string?
		}

		return fmt.Errorf("Failed to refresh access token: %v", err)
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
	var response exchangeTokenResponseModel

	url := "https://www.strava.com/oauth/token" +
		"?client_id=" + getStravaClientId() +
		"&client_secret=" + getStravaClientSecret() +
		"&code=" + code +
		"&grant_type=authorization_code"

	err := utils.RelaxedHttpRequest(utils.RelaxedHttpRequestOptions{
		URL:      url,
		Method:   "POST",
		Response: &response,
	})

	if err != nil {
		if err.Error() == "401" {
			return fmt.Errorf("401") // TODO? make this return directly instead of passing outside as string?
		}

		return fmt.Errorf("Failed to exchange code for token: %v", err)
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
