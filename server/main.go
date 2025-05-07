package main

import (
	"fmt"
	"net/http"
	"os"

	"smartmirror.server/config"
	"smartmirror.server/routes"
)

func main() {
	config.ValidateEnvVars()

	serverPort := os.Getenv(config.EnvServerPort)

	fmt.Printf("Starting the application on port %s\n", serverPort)

	router := routes.SetupRouter()

	http.ListenAndServe(":"+serverPort, router)
}
