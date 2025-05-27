package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"smartmirror.server/version"
)

func RegisterVersionHashRoutes(router *chi.Mux) {
	router.HandleFunc("/version-hash", versionHashHandler)
}

type versionHashHandlerResponse struct {
	Hash string `json:"versionHash"`
}

func versionHashHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	versionHash := versionHashHandlerResponse{
		Hash: version.GetVersion(),
	}

	if err := json.NewEncoder(res).Encode(versionHash); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
