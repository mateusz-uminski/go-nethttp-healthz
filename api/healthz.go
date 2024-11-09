package api

import (
	"encoding/json"
	"net/http"

	"github.com/mateusz-uminski/go-nethttp-healthz/util/log"
)

type Handler func(http.ResponseWriter, *http.Request)

func Healthz(l log.Logger) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		response := HealthzResponse{
			Status: "healthy",
		}

		l.Info("Healthz endpoint requested", "response_status", response.Status)

		jsonResponse(w, http.StatusOK, response) // nolint:errcheck
	}
}

type HealthzResponse struct {
	Status string `json:"status"`
}

func jsonResponse[T any](w http.ResponseWriter, status int, data T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
