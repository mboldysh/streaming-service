package router

import (
	"net/http"
)

type Router interface {
	Handler() http.Handler
	Path() string
}

type Endpoint interface {
	Router
	Method() string
}
