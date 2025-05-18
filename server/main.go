package main

import (
	"fmt"
	"net/http"
	"os"

	"smartmirror.server/config"
	"smartmirror.server/routes"
)

func main() {
	config.SetAndValidateEnvVars()

	serverPort := os.Getenv(config.EnvServerPort)

	fmt.Printf("Starting the application on http://localhost:%s\n", serverPort)

	router := routes.SetupRouter()

	http.ListenAndServe(":"+serverPort, router)
}
