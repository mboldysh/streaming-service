package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mboldysh/streaming-service/pkg/router"
)

type Server struct {
	listenAddr string
	handler    *chi.Mux
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		handler:    chi.NewRouter(),
	}
}

func (s *Server) InitRoutes(routes ...router.Router) {
	s.handler.Route("/api/v1", func(r chi.Router) {
		for _, route := range routes {
			r.Mount(route.Prefix(), route.Handler())
		}
	})
}

func (s *Server) InitMiddleware(middlewares ...func(http.Handler) http.Handler) {
	s.handler.Use(middlewares...)
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    s.listenAddr,
		Handler: s.handler,
	}
	log.Fatal(httpServer.ListenAndServe())
}
