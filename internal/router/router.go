package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Init(routes ...http.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
	)

	r.Route("/api/v1", func(r chi.Router) {
		for _, route := range routes {
			r.Mount("/", route)
		}
	})

	return r
}
