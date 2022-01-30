package service

import (
	"github.com/tzmfreedom/go-sample/mvc/model"
	"github.com/tzmfreedom/go-sample/mvc/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) UpdateAndNotify(id model.UserID) (*model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	err = s.repo.Update(user)
	if err != nil {
		return nil, err
	}
	// Notify
	return nil, err
}
