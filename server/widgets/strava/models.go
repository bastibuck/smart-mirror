package strava

// annual
type annualStatsModel struct {
	Running sportStatsModel `json:"running"`
	Cycling sportStatsModel `json:"cycling"`
	Hiking  sportStatsModel `json:"hiking"`
	Kiting  sportStatsModel `json:"kiting"`
}

type sportStatsModel struct {
	Count       int `json:"count"`
	MovingTimeS int `json:"moving_time_s"`
	DistanceM   int `json:"distance_m"`
}

type credentialsModel struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int    `json:"expires_at"`
}

// last activity
type lastActivityModel struct {
	Name        string      `json:"name"`
	Date        string      `json:"date"`
	Type        string      `json:"type"`
	DistanceM   int         `json:"distance_m"`
	MovingTimeS int         `json:"moving_time_s"`
	Coordinates [][]float64 `json:"coordinates"`
}
