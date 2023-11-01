package api

import (
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) error {
	return WriteJson(w, http.StatusOK, HealthResponse{Status: "UP"})
}
