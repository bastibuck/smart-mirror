package strava

type stravaStats struct {
	Running sportStats `json:"running"`
	Cycling sportStats `json:"cycling"`
}

type sportStats struct {
	Count       int `json:"count"`
	MovingTimeS int `json:"moving_time_s"`
	DistanceM   int `json:"distance_m"`
}
