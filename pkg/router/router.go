package router

import (
	"net/http"
)

type Router interface {
	Handler() http.Handler
	Prefix() string
	ServeHttp(w http.ResponseWriter, r *http.Request)
}

type localRouter struct {
	handler http.Handler
	prefix  string
}

func (r localRouter) Handler() http.Handler {
	return r.handler
}

func (r localRouter) Prefix() string {
	return r.prefix
}

func New(prefix string, handler http.Handler) Router {
	return localRouter{handler, prefix}
}

func (router localRouter) ServeHttp(w http.ResponseWriter, r *http.Request) {
	router.handler.ServeHTTP(w, r)
}
