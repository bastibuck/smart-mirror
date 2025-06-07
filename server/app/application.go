package app

import (
	"fmt"
	"net/http"

	"smartmirror.server/router"
	"smartmirror.server/strava"
	"smartmirror.server/version"
)

func SetupApp() {
	setupAppEnv()

	router := router.SetupRouter([]string{getCorsAllowedOrigin()})

	RegisterWidgets([]Widget{
		version.NewVersionWidget(),
		strava.NewStravaWidget(),
	}, router)

	// Start the server
	serverPort := getServerPort()
	fmt.Printf("Starting the application on port %s\n", serverPort)

	http.ListenAndServe(":"+serverPort, router)
}
