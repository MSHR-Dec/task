package input

import (
	"github.com/MSHR-Dec/task/go_backend/internal/usecase/dto"
)

type UserInputPort interface {
	SignUp(input dto.SignUpInput) error
	SignIn(input dto.SignInInput) (dto.SignInOutput, error)
}
