package strava

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
