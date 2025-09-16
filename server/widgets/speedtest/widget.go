package speedtest

import (
	"time"

	"github.com/go-chi/chi/v5"
	"smartmirror.server/utils"
	"smartmirror.server/widgets"
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

	cron.Schedule("runSpeedTest", 5*time.Minute, func() {
		speedTestResponse, err := runSpeedtest()

		if err != nil {
			logger.Info("Error running speedtest: %v", err)
			speedtestHistory = append(speedtestHistory, SpeedtestHistory{
				Time: time.Now(),
			})
			return
		}

		logger.Info("Speedtest result: %+v", speedTestResponse)

		speedtestHistory = append(speedtestHistory, SpeedtestHistory{
			Time:              time.Now(),
			speedtestResponse: speedTestResponse,
		})
	})

	return &SpeedtestWidget{}
}
