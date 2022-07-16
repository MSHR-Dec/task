package service

import (
	"github.com/MSHR-Dec/task/go_backend/internal/domain/entity"
	"github.com/MSHR-Dec/task/go_backend/internal/domain/repository"
	"github.com/MSHR-Dec/task/go_backend/pkg/oops"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return UserService{
		userRepo: userRepo,
	}
}

func (s UserService) Exist(user entity.User) (bool, error) {
	_, err := s.userRepo.FindByName(user.Name)

	switch err.(type) {
	case oops.NotFound:
		return false, nil
	case oops.InternalServerError:
		return false, err
	default:
		return true, nil
	}
}
