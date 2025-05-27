package main

import (
	"fmt"
	"net/http"
	"os"

	"smartmirror.server/env"
	"smartmirror.server/router"
)

func main() {
	env.SetAndValidateEnv()

	router := router.SetupRouter()

	serverPort := os.Getenv(env.EnvServerPort)
	fmt.Printf("Starting the application on port %s\n", serverPort)

	http.ListenAndServe(":"+serverPort, router)
}
