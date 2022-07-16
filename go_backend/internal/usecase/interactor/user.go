package interactor

import (
	"time"

	"github.com/MSHR-Dec/task/go_backend/internal/domain/factory"
	"github.com/MSHR-Dec/task/go_backend/internal/domain/repository"
	"github.com/MSHR-Dec/task/go_backend/internal/domain/service"
	"github.com/MSHR-Dec/task/go_backend/internal/domain/vo"
	"github.com/MSHR-Dec/task/go_backend/internal/usecase/dto"
	"github.com/MSHR-Dec/task/go_backend/internal/usecase/input"
	"github.com/MSHR-Dec/task/go_backend/pkg/oops"
)

type UserInteractor struct {
	userFactory factory.UserFactory
	userRepo    repository.UserRepository
	userService service.UserService
}

func NewUserInteractor(userFactory factory.UserFactory,
	useRepo repository.UserRepository,
	userService service.UserService,
) input.UserInputPort {
	return UserInteractor{
		userFactory: userFactory,
		userRepo:    useRepo,
		userService: userService,
	}
}

func (i UserInteractor) SignUp(input dto.SignUpInput) error {
	now := time.Now()
	user, err := i.userFactory.Create(input.Name, input.Password, now)
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

	if err = i.userRepo.Save(user, now); err != nil {
		return err
	}

	return nil
}

func (i UserInteractor) SignIn(input dto.SignInInput) (dto.SignInOutput, error) {
	userName, err := vo.NewUserName(input.Name)
	if err != nil {
		return dto.SignInOutput{}, err
	}

	user, err := i.userRepo.FindByName(userName)
	if err != nil {
		return dto.SignInOutput{}, err
	}

	if !user.Password.IsSame(input.Password) {
		return dto.SignInOutput{}, oops.BadRequest{Message: "incorrect password"}
	}

	if user.ShouldUpdatePassword() {
		return dto.SignInOutput{ShouldUpdatePassword: true}, nil
	}

	return dto.SignInOutput{ShouldUpdatePassword: false}, nil
}
