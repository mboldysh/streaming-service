package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mboldysh/streaming-service/internal/router/healthcheck"
)

func Init(routes ...http.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
	)

	r.Get("/api/v1/health", healthcheck.HealthHandler)

	r.Route("/api/v1", func(r chi.Router) {
		for _, route := range routes {
			r.Mount("/", route)
		}
	})

	return r
}
