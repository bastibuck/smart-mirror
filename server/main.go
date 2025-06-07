package main

import (
	"fmt"
	"net/http"

	"smartmirror.server/env"
	"smartmirror.server/router"
	"smartmirror.server/shared"
)

func main() {
	env.SetupEnv()

	router := router.SetupRouter()

	serverPort := shared.GetServerPort()
	fmt.Printf("Starting the application on port %s\n", serverPort)

	http.ListenAndServe(":"+serverPort, router)
}
