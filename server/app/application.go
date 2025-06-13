package app

import (
	"fmt"
	"net/http"

	"smartmirror.server/router"

	"smartmirror.server/widgets"
	"smartmirror.server/widgets/garmin"
	"smartmirror.server/widgets/shared"
	"smartmirror.server/widgets/strava"
	"smartmirror.server/widgets/version"
)

func SetupApp() {
	setupAppEnv()

	router := router.SetupRouter()

	widgets.RegisterWidgets([]widgets.Widget{
		shared.NewSharedWidget(),
		version.NewVersionWidget(),
		strava.NewStravaWidget(),
		garmin.NewGarminWidget(),
	}, router)

	// Start the server
	serverPort := getServerPort()
	fmt.Printf("Starting the application on port %s\n", serverPort)

	http.ListenAndServe(":"+serverPort, router)
}
