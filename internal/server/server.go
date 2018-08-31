package server

import (
	"log"
	"net/http"
)

type Server struct {
	listenAddr string
	handler    http.Handler
}

func NewServer(listenAddr string, handler http.Handler) *Server {
	return &Server{
		listenAddr: listenAddr,
		handler:    handler,
	}
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    s.listenAddr,
		Handler: s.handler,
	}
	log.Fatal(httpServer.ListenAndServe())
}
