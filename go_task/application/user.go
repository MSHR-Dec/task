package application

import (
	"time"

	"github.com/MSHR-Dec/MSHR-Doc/mypkg/oops"

	"github.com/MSHR-Dec/task/go_task/domain/model"
	"github.com/MSHR-Dec/task/go_task/domain/repository"
	"github.com/MSHR-Dec/task/go_task/domain/service"
)

type UserApplication interface {
	SignUp(input SignUpInput) error
}

type UserInteractor struct {
	userRepository repository.UserRepository
	userService    service.UserService
}

func NewUserInteractor(userRepository repository.UserRepository, userService service.UserService) UserInteractor {
	return UserInteractor{
		userRepository: userRepository,
		userService:    userService,
	}
}

type SignUpInput struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (i UserInteractor) SignUp(input SignUpInput) error {
	now := time.Now()
	user, err := model.NewUser(input.Name, input.Password, now)
	if err != nil {
		return err
	}

	ok, err := i.userService.Exist(user)
	if ok {
		return oops.BadRequest{Message: "already exist"}
	}
	if err != nil {
		return err
	}

	if err = i.userRepository.Save(user); err != nil {
		return err
	}

	return nil
}

type SignInInput struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type SignInOutput struct {
	ShouldUpdatePassword bool `json:"should_update_password"`
}

func (i UserInteractor) SignIn(input SignInInput) (SignInOutput, error) {
	name, err := model.NewUserName(input.Name)
	if err != nil {
		return SignInOutput{}, err
	}

	user, err := i.userRepository.FindByName(name)
	if err != nil {
		return SignInOutput{}, err
	}

	if !user.Password.IsSame(input.Password) {
		return SignInOutput{}, oops.BadRequest{Message: "incorrect password"}
	}

	return SignInOutput{ShouldUpdatePassword: user.ShouldUpdatePassword()}, nil
}
