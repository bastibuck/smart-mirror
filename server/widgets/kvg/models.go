package kvg

type departure struct {
	Line          string `json:"line"`
	Destination   string `json:"destination"`
	DepartureTime string `json:"departureTime"`
	DelayMinutes  int    `json:"delayMinutes,omitempty"`
}

type nextDeparturesResponse struct {
	StopName   string      `json:"stopName"`
	Departures []departure `json:"departures"`
}
