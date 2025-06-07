package version

import (
	"encoding/json"
	"net/http"
)

type versionHashHandlerResponse struct {
	Hash string `json:"versionHash"`
}

func versionHashHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	versionHash := versionHashHandlerResponse{
		Hash: getVersionHash(),
	}

	if err := json.NewEncoder(res).Encode(versionHash); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
