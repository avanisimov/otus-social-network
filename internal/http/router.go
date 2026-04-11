package http

import "github.com/go-chi/chi/v5"

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	h := NewHandler()

	r.Post("/user/register", h.Register)
	r.Get("/user/get/{id}", h.GetUser)
	r.Post("/login", h.Login)

	return r
}
