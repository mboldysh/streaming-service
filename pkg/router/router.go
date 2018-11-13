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

func (router localRouter) ServeHttp(w http.ResponseWriter, r *http.Request) {
	router.handler.ServeHTTP(w, r)
}

//NewRouter creates a single router
func NewRouter(pattern string, handlerFunc http.HandlerFunc) Router {
	return localRouter{handlerFunc, pattern}
}

//NewHandler creates a handler
func NewHandler(prefix string, handler http.Handler) Router {
	return localRouter{handler, prefix}
}
