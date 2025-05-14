package tasks

import (
	"step1_simple_api/internal/db"
	"step1_simple_api/internal/types"
)

type Service struct {
	db db.DB
}

func New(db db.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) ObjectsTasks() ([]types.Task, error) {
	return s.db.ObjectsTasks()
}

func (s *Service) ObjectTask(objectID string) (types.Task, error) {
	return s.db.ObjectTask(objectID)
}

func (s *Service) CreateTask(task types.Task) (types.Task, error) {
	return s.db.CreateTask(task)
}

func (s *Service) UpdateTask(task types.Task) error {
	return s.db.UpdateTask(task)
}

func (s *Service) DestroyTask(objectID string) error {
	return s.db.DestroyTask(objectID)
}
