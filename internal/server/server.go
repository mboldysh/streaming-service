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

//New initialize a new Server.
func New(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		handler:    chi.NewRouter(),
	}
}

//InitRoutes initializes the list of routers for the server.
func (s *Server) InitRoutes(routes ...router.Router) {
	s.handler.Route("/api/v1", func(r chi.Router) {
		for _, route := range routes {
			endpoint, ok := route.(router.Endpoint)
			if ok {
				r.Method(endpoint.Method(), endpoint.Path(), endpoint.Handler())
			} else {
				r.Mount(route.Path(), route.Handler())
			}
		}
	})
}

//InitMiddleware initialize the list of middlewares for the server
func (s *Server) InitMiddleware(middlewares ...func(http.Handler) http.Handler) {
	s.handler.Use(middlewares...)
}

//Run run the server
func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    s.listenAddr,
		Handler: s.handler,
	}
	log.Fatal(httpServer.ListenAndServe())
}
