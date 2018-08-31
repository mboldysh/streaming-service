package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func Init(routes ...http.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
	)

	r.Route("/api/v1", func(r chi.Router) {
		for _, route := range routes {
			r.Mount("/", route)
		}
	})

	return r
}
