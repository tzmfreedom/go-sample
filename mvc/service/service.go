package service

import (
	"github.com/tzmfreedom/go-sample/mvc/model"
	"github.com/tzmfreedom/go-sample/mvc/repository"
)

type UserService struct {
	userRepo   *repository.UserRepository
	notifyRepo *repository.UserNotifyRepository
}

func NewUserService(userRepo *repository.UserRepository, notifyRepo *repository.UserNotifyRepository) *UserService {
	return &UserService{
		userRepo:   userRepo,
		notifyRepo: notifyRepo,
	}
}

func (s *UserService) UpdateAndNotify(id model.UserID) (*model.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	err = s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}
	return user, s.notifyRepo.Notify(user)
}
