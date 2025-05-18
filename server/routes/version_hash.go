package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"smartmirror.server/config"
)

type VersionHashResponse struct {
	Hash string `json:"versionHash"`
}

var VersionHash string

func VersionHashHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	versionHash := VersionHashResponse{
		Hash: os.Getenv(config.EnvVersionHash),
	}

	if err := json.NewEncoder(res).Encode(versionHash); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
