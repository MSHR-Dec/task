package service

import (
	"github.com/MSHR-Dec/MSHR-Doc/mypkg/oops"

	"github.com/MSHR-Dec/task/go_task/domain/model"
	"github.com/MSHR-Dec/task/go_task/domain/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return UserService{
		userRepository: userRepo,
	}
}

func (s UserService) Exist(user model.User) (bool, error) {
	_, err := s.userRepository.FindByName(user.Name)

	switch err.(type) {
	case oops.NotFound:
		return false, nil
	case oops.InternalServerError:
		return false, err
	default:
		return true, nil
	}
}
