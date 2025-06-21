package kvg

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	kvgStopInfoUrl := fmt.Sprintf("https://kvg-internetservice-proxy.p.networkteam.com/internetservice/services/passageInfo/stopPassages/stop?stop=%s", getHomeStopID())

	req, err := http.NewRequest("GET", kvgStopInfoUrl, nil)
	if err != nil {
		return nextDeparturesResponse{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nextDeparturesResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nextDeparturesResponse{}, fmt.Errorf("KVG API returned status: %s", resp.Status)
	}

	var response kvgStopInfoResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nextDeparturesResponse{}, err
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
