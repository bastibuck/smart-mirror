package garmin

import (
	"fmt"
	"time"

	"github.com/bastibuck/go-garmin"
)

func getSevenDaySteps() (sevenDayStepsModel, error) {
	if cachedData, found := garminCache.getSevenDaySteps(); found {
		return cachedData, nil
	}

	client := garmin.NewClient()
	err := client.Login(getEmail(), getPassword())

	if err != nil {
		return sevenDayStepsModel{}, fmt.Errorf("failed to login to Garmin: %w", err)
	}

	api := garmin.NewAPI(client)

	today := time.Now()

	steps, err := api.UserSummary.DailySteps(
		today.AddDate(0, 0, -6),
		today,
	)

	if err != nil {
		return sevenDayStepsModel{}, fmt.Errorf("failed to get daily steps: %w", err)
	}

	total := 0
	dayCount := len(steps.Values)
	days := make([]daySteps, 0, dayCount)

	for _, day := range steps.Values {
		total += day.Values.TotalSteps

		days = append(days, daySteps{
			Steps: day.Values.TotalSteps,
			Date:  day.CalendarDate,
		})
	}

	result := sevenDayStepsModel{
		Average: int(total / dayCount),
		Total:   total,
		Days:    days,
	}

	garminCache.setSevenDaySteps(result)

	return result, nil
}
