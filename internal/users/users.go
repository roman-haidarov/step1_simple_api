package users

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

func (s *Service) GetUsers() ([]types.User, error) {
	return s.db.GetUsers()
}

func (s *Service) GetUser(objectID int) (types.User, error) {
	return s.db.GetUser(objectID)
}

func (s *Service) CreateUser(user types.User) (types.User, error) {
	return s.db.CreateUser(user)
}

func (s *Service) UpdateUser(user types.User) error {
	return s.db.UpdateUser(user)
}

func (s *Service) DestroyUser(objectID int) error {
	return s.db.DestroyUser(objectID)
}
