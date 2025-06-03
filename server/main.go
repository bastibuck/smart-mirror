package main

import (
	"fmt"
	"net/http"

	"smartmirror.server/env"
	"smartmirror.server/router"
	"smartmirror.server/shared"
	// "github.com/twpayne/go-polyline"
)

func main() {
	env.SetupEnv()

	// polylineString := "_p~iF~ps|U_ulLnnqC_mqNvxq`@"
	// buf := []byte(polylineString)
	// coords, _, _ := polyline.DecodeCoords(buf)
	// fmt.Println(coords)

	router := router.SetupRouter()

	serverPort := shared.GetServerPort()
	fmt.Printf("Starting the application on port %s\n", serverPort)

	http.ListenAndServe(":"+serverPort, router)
}
