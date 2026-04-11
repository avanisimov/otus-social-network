package http

import (
	"encoding/json"
	"net/http"

	"github.com/avanisimov/otus-social-network/internal/user"
)

type Handler struct {
	userService *user.Service
}

func NewHandler(us *user.Service) *Handler {
	return &Handler{userService: us}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req user.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.userService.Register(r.Context(), req)
	if err != nil {
		http.Error(w, "cannot register", 500)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok - get user"}`))
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok - login"}`))
}
