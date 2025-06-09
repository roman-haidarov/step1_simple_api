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

func (s *Service) GetTasks(userID ...int) ([]types.Task, error) {
	return s.db.GetTasks(userID...)
}

func (s *Service) GetTask(objectID string) (types.Task, error) {
	return s.db.GetTask(objectID)
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
