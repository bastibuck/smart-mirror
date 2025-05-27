package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"smartmirror.server/env"
)

type VersionHashResponse struct {
	Hash string `json:"versionHash"`
}

func RegisterVersionHashRoutes(router *chi.Mux) {
	router.HandleFunc("/version-hash", versionHashHandler)
}

func versionHashHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	versionHash := VersionHashResponse{
		Hash: os.Getenv(env.EnvVersionHash),
	}

	if err := json.NewEncoder(res).Encode(versionHash); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
