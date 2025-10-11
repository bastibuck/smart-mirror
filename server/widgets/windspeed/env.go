package windspeed

import (
	"os"
	"strconv"
)

const (
	envLong = "WINDSPEED_LONGITUDE"
	envLat  = "WINDSPEED_LATITUDE"
)

func getEnvKeys() []string {
	return []string{
		envLong,
		envLat,
	}
}

func getGpsCoordinates() struct {
	Longitude float64
	Latitude  float64
} {
	long, _ := strconv.ParseFloat(os.Getenv(envLong), 64)
	lat, _ := strconv.ParseFloat(os.Getenv(envLat), 64)

	return struct {
		Longitude float64
		Latitude  float64
	}{
		Longitude: long,
		Latitude:  lat,
	}
}
