package speedtest

import (
	"math/rand"
	"time"

	"github.com/go-chi/chi/v5"
	"smartmirror.server/utils"
	"smartmirror.server/widgets"
	"smartmirror.server/widgets/shared"
)

type SpeedtestWidget struct{}

type SpeedtestHistory struct {
	speedtestResponse
	Time time.Time
}

var speedtestHistory []SpeedtestHistory

var _ widgets.Widget = (*SpeedtestWidget)(nil)

func (v *SpeedtestWidget) SetupEnv() {
}

func (v *SpeedtestWidget) SetupRouter(router *chi.Mux) {
	router.HandleFunc("/speedtest", speedtestHandler)
}

func NewSpeedtestWidget() *SpeedtestWidget {
	cron := utils.NewCron("SPEEDTEST")

	// use fixed data in development mode
	if shared.GetAppMode() == "development" {
		logger.Info("Using fixed speedtest data in development mode")

		entryCount := 12 // 1 hour of data (5 minute intervals)
		for i := entryCount; i > 0; i-- {
			speedtestHistory = append(speedtestHistory, SpeedtestHistory{
				Time: time.Now().Add(-time.Duration(i*5) * time.Minute),
				speedtestResponse: speedtestResponse{
					Download: float64(100 + rand.Intn(10) + i),
					Upload:   float64(50 + rand.Intn(10) + i),
					Ping:     int64(20 + rand.Intn(5) + i),
				},
			})
		}

		return &SpeedtestWidget{}
	}

	cron.Schedule("runSpeedTest", 5*time.Minute, func() {
		speedTestResponse, err := runSpeedtest()

		if err != nil {
			logger.Info("Error running speedtest: %v", err)
			addSpeedtestResultToHistory(speedtestResponse{}) // add empty response
			return
		}

		logger.Info("Speedtest result: %+v", speedTestResponse)

		addSpeedtestResultToHistory(speedTestResponse)
	})

	return &SpeedtestWidget{}
}

func addSpeedtestResultToHistory(result speedtestResponse) {
	// always prepend new result so we can keep the history in reverse chronological order to break early on a GET
	speedtestHistory = append([]SpeedtestHistory{
		{
			Time:              time.Now(),
			speedtestResponse: result,
		},
	}, speedtestHistory...)
}
