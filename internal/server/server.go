package server

import "step1_simple_api/internal/api"

type Server struct {
	api *api.API
}

func New() (*Server, error) {
	s := &Server{}
	s.api = api.New()

	return s, nil
}

func (s *Server) Run() error {
	return s.api.Serve()
}
