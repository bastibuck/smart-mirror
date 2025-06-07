package strava

// annual
type stravaStats struct {
	Running sportStats `json:"running"`
	Cycling sportStats `json:"cycling"`
	Hiking  sportStats `json:"hiking"`
	Kiting  sportStats `json:"kiting"`
}

type sportStats struct {
	Count       int `json:"count"`
	MovingTimeS int `json:"moving_time_s"`
	DistanceM   int `json:"distance_m"`
}

type credentials struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int    `json:"expires_at"`
}

// last activity
type lastActivity struct {
	Type        string      `json:"type"`
	DistanceM   int         `json:"distance_m"`
	MovingTimeS int         `json:"moving_time_s"`
	Coordinates [][]float64 `json:"coordinates"`
}
