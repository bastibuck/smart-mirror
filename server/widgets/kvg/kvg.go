package kvg

import (
	"fmt"
	"time"

	"smartmirror.server/utils"
)

type kvgStopInfoResponse struct {
	StopName string `json:"stopName"`
	Actual   []struct {
		Line        string `json:"patternText"`
		Destination string `json:"direction"`
		PlannedTime string `json:"plannedTime"`
		ActualTime  string `json:"actualTime"`
	} `json:"actual"`
	GeneralAlerts []struct {
		Title string `json:"title"`
	} `json:"generalAlerts"`
}

func fetchNextDepartures(limit int) (nextDeparturesResponse, error) {
	var response kvgStopInfoResponse

	err := utils.RelaxedHttpRequest(utils.RelaxedHttpRequestOptions{
		URL:      fmt.Sprintf("https://kvg-internetservice-proxy.p.networkteam.com/internetservice/services/passageInfo/stopPassages/stop?stop=%s", getHomeStopID()),
		Response: &response,
		Delay: utils.RelaxedHttpRequestDelay{
			Variance: 50,
			Average:  1000,
		},
		Timeout: time.Second * 10,
		Retries: 3,
	})
	if err != nil {
		return nextDeparturesResponse{}, fmt.Errorf("Failed to fetch KVG stop info: %v", err)
	}

	departures := make([]departure, 0, len(response.Actual))
	for _, dep := range response.Actual {

		// use planned time if actual time is empty
		actualTime := dep.ActualTime
		if actualTime == "" {
			logger.Info("Falling back to planned time %s for line %s", dep.PlannedTime, dep.Line)
			actualTime = dep.PlannedTime
		}

		delay, err := utils.MinutesBetween(dep.PlannedTime, actualTime)

		if err != nil {
			logger.Info("Failed to calculate delay for departure %s: %v", dep.Line, err)
			delay = 0 // If we can't calculate the delay, default to 0
		}

		// Ensure delay is non-negative
		if delay < 0 {
			delay = 0 // If the delay is negative, set it to 0
		}

		departures = append(departures, departure{
			Line:          dep.Line,
			Destination:   dep.Destination,
			DepartureTime: actualTime,
			DelayMinutes:  delay,
		})
	}

	departuresLimit := limit
	if len(departures) < departuresLimit {
		departuresLimit = len(departures)
	}

	// General alerts
	alerts := make([]string, 0, len(response.GeneralAlerts))
	for _, alert := range response.GeneralAlerts {
		alerts = append(alerts, alert.Title)
	}

	return nextDeparturesResponse{
		StopName:   response.StopName,
		Departures: departures[:departuresLimit],
		Alerts:     alerts,
	}, nil
}
