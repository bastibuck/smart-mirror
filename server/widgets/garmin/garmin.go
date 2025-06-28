package garmin

import (
	"fmt"
	"time"

	"github.com/bastibuck/go-garmin"
)

var GET_STEPS_ERROR_COUNT int = 0

func getSevenDaySteps(apiClient *garmin.API) (sevenDayStepsModel, error) {
	if cachedData, found := garminCache.getSevenDaySteps(); found {
		logger.Info("Using cached seven day steps data %d", cachedData.Total)
		return cachedData, nil
	}

	logger.Info("Step error count: %d", GET_STEPS_ERROR_COUNT)

	today := time.Now()

	if GET_STEPS_ERROR_COUNT >= 3 {
		return sevenDayStepsModel{}, fmt.Errorf("too many requests for daily steps, please try again later")
	}

	steps, err := apiClient.UserSummary.DailySteps(
		today.AddDate(0, 0, -6),
		today,
	)

	if err != nil {
		GET_STEPS_ERROR_COUNT++
		return sevenDayStepsModel{}, fmt.Errorf("failed to get daily steps: %w", err)
	}

	if GET_STEPS_ERROR_COUNT > 0 {
		logger.Info("Decreasing Step error count by 1 due to successful request")
		GET_STEPS_ERROR_COUNT--
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

	logger.Info("Fetched steps: %d", total)

	garminCache.setSevenDaySteps(result)

	return result, nil
}
