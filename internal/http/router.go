package http

import (
	"net/http"

	"github.com/avanisimov/otus-social-network/internal/user"

	"github.com/go-chi/chi/v5"
)

func NewRouter(userService *user.Service, db DB) *chi.Mux {
	r := chi.NewRouter()

	// Middleware to set Content-Type to application/json for all routes
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	h := NewHandler(userService)
	healthHandler := NewHealthHandler(db)

	r.Get("/healthz", healthHandler.Healthz)
	r.Get("/readiness", healthHandler.Readiness)
	r.Handle("/metrics", healthHandler.MetricsHandler())

	r.Post("/user/register", h.Register)
	r.Get("/user/get/{id}", h.GetUser)
	r.Get("/users/search", h.SearchUsers)
	r.Post("/login", h.Login)

	return r
}
