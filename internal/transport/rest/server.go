package rest

import (
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(addr string, handler http.Handler) *Server {

	return &Server{
		server: &http.Server{
			Addr:    ":" + addr,
			Handler: handler,
		},
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.server.Close()
}
