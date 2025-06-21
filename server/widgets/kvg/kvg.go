package kvg

import (
	"fmt"

	"smartmirror.server/utils"
)

type kvgStopInfoResponse struct {
	StopName string `json:"stopName"`
	Actual   []struct {
		Line        string `json:"patternText"`
		Destination string `json:"direction"`
		PlannedTime string `json:"plannedTime"`
		ActualTime  string `json:"actualTime"`
	}
}

func fetchNextDepartures(limit int) (nextDeparturesResponse, error) {
	var response kvgStopInfoResponse

	err := utils.RelaxedHttpRequest(utils.RelaxedHttpRequestOptions{
		URL:      fmt.Sprintf("https://kvg-internetservice-proxy.p.networkteam.com/internetservice/services/passageInfo/stopPassages/stop?stop=%s", getHomeStopID()),
		Response: &response,
	})
	if err != nil {
		return nextDeparturesResponse{}, fmt.Errorf("Failed to fetch KVG stop info: %v", err)
	}

	var departures []departure
	for _, dep := range response.Actual {

		delay, err := utils.MinutesBetween(dep.PlannedTime, dep.ActualTime)

		if err != nil {
			fmt.Print(fmt.Errorf("Failed to calculate delay for departure %s: %v", dep.Line, err))
			delay = 0 // If we can't calculate the delay, default to 0
		}

		// Ensure delay is non-negative
		if delay < 0 {
			delay = 0 // If the delay is negative, set it to 0
		}

		departures = append(departures, departure{
			Line:          dep.Line,
			Destination:   dep.Destination,
			DepartureTime: dep.ActualTime,
			DelayMinutes:  delay,
		})
	}

	departuresLimit := limit
	if len(departures) < departuresLimit {
		departuresLimit = len(departures)
	}

	return nextDeparturesResponse{
		StopName:   response.StopName,
		Departures: departures[:departuresLimit],
	}, nil
}
