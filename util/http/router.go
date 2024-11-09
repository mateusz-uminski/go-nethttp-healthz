package http

import (
	"fmt"
	"net/http"
)

type Endpoint struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
}

type router struct {
	mux *http.ServeMux
}

func NewRouter() *router {
	return &router{
		mux: http.NewServeMux(),
	}
}

func (r *router) ServeMux() *http.ServeMux {
	return r.mux
}

func (r *router) RegisterEndpoint(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.mux.HandleFunc(path, handler)
}

func (r *router) RegisterEndpoints(prefix string, endpoints []Endpoint) {
	for _, ep := range endpoints {
		path := fmt.Sprintf("%s%s", prefix, ep.Path)
		r.RegisterEndpoint(path, ep.Handler)
	}
}
