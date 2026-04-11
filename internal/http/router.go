package http

import (
	"github.com/avanisimov/otus-social-network/internal/user"

	"github.com/go-chi/chi/v5"
)

func NewRouter(userService *user.Service) *chi.Mux {
	r := chi.NewRouter()

	h := NewHandler(userService)

	r.Post("/user/register", h.Register)
	r.Get("/user/get/{id}", h.GetUser)
	r.Post("/login", h.Login)

	return r
}
