package windspeed

type windspeedModel struct {
	WindSpeedKn     float64 `json:"wind_speed_kn"`
	GustSpeedKn     float64 `json:"gust_speed_kn"`
	WindDirectionKn int16   `json:"wind_direction_kn"`
}
