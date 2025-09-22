package app

import (
	"fmt"
	"net/http"

	"smartmirror.server/router"

	"smartmirror.server/widgets"
	"smartmirror.server/widgets/kptncook"
	"smartmirror.server/widgets/kvg"
	"smartmirror.server/widgets/shared"
	"smartmirror.server/widgets/speedtest"
	"smartmirror.server/widgets/strava"
	"smartmirror.server/widgets/version"
	"smartmirror.server/widgets/windspeed"
)

func SetupApp() {
	setupAppEnv()

	router := router.SetupRouter()

	widgets.RegisterWidgets([]widgets.Widget{
		shared.NewSharedWidget(),
		version.NewVersionWidget(),
		strava.NewStravaWidget(),
		kvg.NewKVGWidget(),
		kptncook.NewKptnCookWidget(),
		speedtest.NewSpeedtestWidget(),
		windspeed.NewWindspeedWidget(),
	}, router)

	// Start the server
	serverPort := getServerPort()
	fmt.Printf("Starting the application on port %s\n", serverPort)

	http.ListenAndServe(":"+serverPort, router)
}
