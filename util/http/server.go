package http

import (
	"fmt"
	"net/http"
	"time"
)

type Router interface {
	RegisterEndpoint(path string, handler func(http.ResponseWriter, *http.Request))
	ServeMux() *http.ServeMux
}

type server struct {
	host   string
	port   int
	router Router
	server *http.Server
}

type serverOption func(s *server)

func NewServer(options ...serverOption) *server {
	s := &server{}

	for _, opt := range options {
		opt(s)
	}

	s.server = &http.Server{
		Addr:    s.addr(),
		Handler: s.router.ServeMux(),
	}

	return s
}

func ServerWithAddr(host string) serverOption {
	return func(s *server) {
		s.host = host
	}
}

func ServerWithPort(port int) serverOption {
	return func(s *server) {
		s.port = port
	}
}

func ServerWithTimeout(timeout time.Duration) serverOption {
	return func(s *server) {
		s.server.ReadTimeout = time.Second * timeout
		s.server.WriteTimeout = time.Second * timeout
	}
}

func ServerWithRouter(router Router) serverOption {
	return func(s *server) {
		s.router = router
	}
}

func (s *server) Start() {
	s.server.ListenAndServe() // nolint:errcheck
}

func (s *server) addr() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}
