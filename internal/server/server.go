package server

import (
	"context"
	"step1_simple_api/internal/api"
	"step1_simple_api/internal/db"
	"step1_simple_api/internal/tasks"

	"github.com/sirupsen/logrus"
)

type Server struct {
	db    db.DB
	api   *api.API
	tasks *tasks.Service
}

func New() (*Server, error) {
	s := Server{}
	db, err := db.InitDB()

	if err != nil {
		return nil, err
	}

	s.db = *db
	s.tasks = tasks.New(s.db)
	s.api = api.New(s.tasks)

	return &s, nil
}

func (s *Server) Run(ctx context.Context) error {
	return s.api.Serve(ctx)
}

func (s *Server) Shutdown() {
	logrus.Info("graceful server shutdown")
}
