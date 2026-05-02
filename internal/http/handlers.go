package http

import (
	"encoding/json"
	"net/http"

	"github.com/avanisimov/otus-social-network/internal/auth"
	"github.com/avanisimov/otus-social-network/internal/user"
	"github.com/go-chi/chi/v5"
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
		http.Error(w, "cannot register"+err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.userService.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req auth.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", 400)
		return
	}

	u, err := h.userService.GetUserWithPassword(r.Context(), req.ID, req.Password)
	if err != nil {
		http.Error(w, "user not found", 404)
		return
	}

	token, _ := auth.GenerateToken(u.ID)
	json.NewEncoder(w).Encode(auth.LoginResponse{Token: token})
}

func (h *Handler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	first_name := r.URL.Query().Get("first_name")
	second_name := r.URL.Query().Get("second_name")
	limit := getIntQuery(r, "limit", 10)
	// For simplicity, we just return all users. In a real app, you'd implement search logic.
	users, err := h.userService.SearchUsers(r.Context(), first_name, second_name, limit)
	if err != nil {
		http.Error(w, "cannot get users", 500)
		return
	}

	json.NewEncoder(w).Encode(UsersSearchResponse{Users: users})
}

type UsersSearchResponse struct {
	Users []user.User `json:"users"`
}
