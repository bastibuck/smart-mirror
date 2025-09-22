package windspeed

import (
	"fmt"

	"smartmirror.server/utils"
)

type currentWindSpeedResponse struct {
	Current struct {
		WindSpeed10m     float64 `json:"wind_speed_10m"`
		WindGusts10m     float64 `json:"wind_gusts_10m"`
		WindDirection10m int16   `json:"wind_direction_10m"`
	} `json:"current"`
}

func getCurrentWind() (windspeedModel, error) {
	var response currentWindSpeedResponse

	coordinates := getGpsCoordinates()

	err := utils.RelaxedHttpRequest(utils.RelaxedHttpRequestOptions{
		URL:      fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current=wind_speed_10m,wind_gusts_10m,wind_direction_10m&wind_speed_unit=kn", coordinates.Latitude, coordinates.Longitude),
		Response: &response,
	})

	if err != nil {
		return windspeedModel{}, fmt.Errorf("Failed to fetch current wind speed data: %v", err)
	}

	return windspeedModel{
		WindSpeedKn:     response.Current.WindSpeed10m,
		GustSpeedKn:     response.Current.WindGusts10m,
		WindDirectionKn: response.Current.WindDirection10m,
	}, nil
}
