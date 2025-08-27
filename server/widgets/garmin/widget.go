package garmin

import (
	"net/http"

	"github.com/bastibuck/go-garmin"
	"github.com/go-chi/chi/v5"

	"smartmirror.server/env"
	"smartmirror.server/widgets"
)

type GarminWidget struct {
	ApiClient *garmin.API
}

var _ widgets.Widget = (*GarminWidget)(nil)

func (v *GarminWidget) SetupEnv() {
	// already done in NewGarminWidget because of login when creating the widget
}

func (v *GarminWidget) SetupRouter(router *chi.Mux) {
	router.HandleFunc("/steps/weekly", func(w http.ResponseWriter, r *http.Request) {
		stepsThisWeekHandler(w, v.ApiClient)
	})
}

func NewGarminWidget() *GarminWidget {
	env.ValidateEnvKeys(getEnvKeys())

	client := garmin.NewClient()
	err := client.Login(getEmail(), getPassword())

	if err != nil {
		panic("failed to login to Garmin")
	}

	api := garmin.NewAPI(client)

	return &GarminWidget{
		ApiClient: api,
	}
}
