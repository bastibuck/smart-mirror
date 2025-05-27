package main

import (
	"fmt"
	"net/http"

	"smartmirror.server/env"
	"smartmirror.server/router"
)

func main() {
	env.SetupEnv()

	router := router.SetupRouter()

	serverPort := env.GetServerPort()
	fmt.Printf("Starting the application on port %s\n", serverPort)

	http.ListenAndServe(":"+serverPort, router)
}
