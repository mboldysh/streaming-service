package router

import (
	"net/http"
)

type localRouter struct {
	path    string
	handler http.Handler
}

type localEndpoint struct {
	method string
	localRouter
}

func (r localRouter) Handler() http.Handler {
	return r.handler
}

func (r localRouter) Path() string {
	return r.path
}

func (e localEndpoint) Method() string {
	return e.method
}

func NewRouter(path string, handler http.Handler) Router {
	return localRouter{path, handler}
}

func NewRouter2(path string, handler func() http.Handler) Router {
	return localRouter{path, handler()}
}

func Get(path string, handler http.HandlerFunc) Endpoint {
	return localEndpoint{http.MethodGet, localRouter{path, handler}}
}

func Head(path string, handler http.HandlerFunc) Endpoint {
	return localEndpoint{http.MethodHead, localRouter{path, handler}}
}

func Post(path string, handler http.HandlerFunc) Endpoint {
	return localEndpoint{http.MethodPost, localRouter{path, handler}}
}

func Put(path string, handler http.HandlerFunc) Endpoint {
	return localEndpoint{http.MethodPut, localRouter{path, handler}}
}

func Delete(path string, handler http.HandlerFunc) Endpoint {
	return localEndpoint{http.MethodDelete, localRouter{path, handler}}
}

func Connect(path string, handler http.HandlerFunc) Endpoint {
	return localEndpoint{http.MethodConnect, localRouter{path, handler}}
}

func Options(path string, handler http.HandlerFunc) Endpoint {
	return localEndpoint{http.MethodOptions, localRouter{path, handler}}
}

func Trace(path string, handler http.HandlerFunc) Endpoint {
	return localEndpoint{http.MethodTrace, localRouter{path, handler}}
}

func Patch(path string, handler http.HandlerFunc) Endpoint {
	return localEndpoint{http.MethodPatch, localRouter{path, handler}}
}
