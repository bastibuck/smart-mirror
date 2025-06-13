package garmin

type sevenDayStepsModel struct {
	Total   int        `json:"total"`
	Average int        `json:"average"`
	Days    []daySteps `json:"days"`
}

type daySteps struct {
	Steps int    `json:"steps"`
	Date  string `json:"date"`
}
