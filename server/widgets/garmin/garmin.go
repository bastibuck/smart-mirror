package garmin

import (
	"fmt"
	"time"

	"github.com/bastibuck/go-garmin"
)

var LOGIN_COUNT int = 0
var GET_STEPS_COUNT int = 0

func getSevenDaySteps() (sevenDayStepsModel, error) {
	if cachedData, found := garminCache.getSevenDaySteps(); found {
		logger("Using cached seven day steps data %d", cachedData.Total)
		return cachedData, nil
	}

	if LOGIN_COUNT >= 3 {
		return sevenDayStepsModel{}, fmt.Errorf("too many login attempts, please try again later")
	}

	client := garmin.NewClient()
	err := client.Login(getEmail(), getPassword())

	if err != nil {
		LOGIN_COUNT++
		logger("Failed to login to garmin: %v", err)
		return sevenDayStepsModel{}, fmt.Errorf("failed to login to Garmin: %w", err)
	}

	api := garmin.NewAPI(client)

	today := time.Now()

	if GET_STEPS_COUNT >= 3 {
		return sevenDayStepsModel{}, fmt.Errorf("too many requests for daily steps, please try again later")
	}

	steps, err := api.UserSummary.DailySteps(
		today.AddDate(0, 0, -6),
		today,
	)

	if err != nil {
		GET_STEPS_COUNT++
		logger("Failed to get daily steps: %v", err)
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

	logger("Fetched steps: %d", total)

	garminCache.setSevenDaySteps(result)

	return result, nil
}
