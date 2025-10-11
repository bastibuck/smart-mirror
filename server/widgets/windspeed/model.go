package windspeed

type windspeedModel struct {
	WindSpeedKn      float64 `json:"windSpeedKn"`
	GustSpeedKn      float64 `json:"gustSpeedKn"`
	WindDirectionDeg int16   `json:"windDirectionDeg"`
}
