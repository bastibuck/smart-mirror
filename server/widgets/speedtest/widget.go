package speedtest

import (
	"github.com/go-chi/chi/v5"
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
	return &SpeedtestWidget{}
}
