package speedtest

import (
	"time"

	"github.com/go-chi/chi/v5"
	"smartmirror.server/utils"
	"smartmirror.server/widgets"
)

type SpeedtestWidget struct{}

var _ widgets.Widget = (*SpeedtestWidget)(nil)

func (v *SpeedtestWidget) SetupEnv() {
}

func (v *SpeedtestWidget) SetupRouter(router *chi.Mux) {
	router.HandleFunc("/speedtest", speedtestHandler)
}

func NewSpeedtestWidget() *SpeedtestWidget {
	cron := utils.NewCron("SPEEDTEST")

	cron.Schedule("runSpeedTest", 15*time.Minute, func() {
		speedTestResponse, err := runSpeedtest()

		if err != nil {
			logger.Info("Error running speedtest: %v", err)
		}

		logger.Info("Speedtest result: %+v", speedTestResponse)
	})

	return &SpeedtestWidget{}
}
