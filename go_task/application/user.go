package application

import (
	"database/sql"
	"time"

	"github.com/MSHR-Dec/MSHR-Doc/mypkg/oops"

	"github.com/MSHR-Dec/task/go_task/application/dto"
	"github.com/MSHR-Dec/task/go_task/domain/model"
	"github.com/MSHR-Dec/task/go_task/domain/repository"
	"github.com/MSHR-Dec/task/go_task/domain/service"
)

type UserApplication interface {
	SignUp(input dto.SignUpInput) (dto.SignUpOutput, error)
	SignIn(input dto.SignInInput) (dto.SignInOutput, error)
	EditProfile(input dto.EditProfileInput) error
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

func (i UserInteractor) SignUp(input dto.SignUpInput) (dto.SignUpOutput, error) {
	now := time.Now()
	user, err := model.NewUser(input.Name, input.Password, now)
	if err != nil {
		return dto.SignUpOutput{}, err
	}

	ok, err := i.userService.Exist(user)
	if ok {
		return dto.SignUpOutput{}, oops.BadRequest{Message: "already exist"}
	}
	if err != nil {
		return dto.SignUpOutput{}, err
	}

	id, err := i.userRepository.Save(user)
	if err != nil {
		return dto.SignUpOutput{}, err
	}

	return dto.SignUpOutput{
		ID: id,
	}, nil
}

func (i UserInteractor) SignIn(input dto.SignInInput) (dto.SignInOutput, error) {
	name, err := model.NewUserName(input.Name)
	if err != nil {
		return dto.SignInOutput{}, err
	}

	user, err := i.userRepository.FindByName(name)
	if err != nil {
		return dto.SignInOutput{}, err
	}

	if !user.Password.IsSame(input.Password) {
		return dto.SignInOutput{}, oops.BadRequest{Message: "incorrect password"}
	}

	return dto.SignInOutput{
		ID:                   int(user.ID),
		ShouldUpdatePassword: user.ShouldUpdatePassword(),
	}, nil
}

func (i UserInteractor) EditProfile(input dto.EditProfileInput) error {
	user, err := i.userRepository.FindByID(uint(input.ID))
	if err != nil {
		return err
	}

	if input.Password != "" {
		if user.Password.IsSame(input.Password) {
			return oops.BadRequest{Message: "same password"}
		}

		password, err := model.NewPassword(input.Password)
		if err != nil {
			return err
		}

		user.Password = password
	}

	now := time.Now()
	user.LastPasswordModifiedAt = now
	user.ModifiedAt = sql.NullTime{
		Time:  now,
		Valid: true,
	}

	if err = i.userRepository.Update(user); err != nil {
		return err
	}

	return nil
}
