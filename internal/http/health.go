package http

import (
	"encoding/json"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type DB interface {
    Ping() error
}

type HealthHandler struct{
	db DB
}

func NewHealthHandler(db DB) *HealthHandler {
	return &HealthHandler{db: db}
}

func (h *HealthHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthzResponse{Status: "ok"})
}

type HealthzResponse struct {
	Status string `json:"status"`
}

func (h *HealthHandler) Readiness(w http.ResponseWriter, r *http.Request) {
	if err := h.db.Ping(); err != nil {
		http.Error(w, "database not ready", http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ReadinessResponse{Status: "ok"})
}

type ReadinessResponse struct {
	Status string `json:"status"`
}

func (h *HealthHandler) MetricsHandler() http.Handler {
    return promhttp.Handler()
}