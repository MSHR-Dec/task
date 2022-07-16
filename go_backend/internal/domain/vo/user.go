package vo

import (
	"unicode/utf8"

	"github.com/google/uuid"

	"github.com/MSHR-Dec/task/go_backend/pkg/oops"
)

type UUID string

func NewUUID() UUID {
	return UUID(uuid.NewString())
}

func (vo UUID) String() string {
	return string(vo)
}

type UserName string

func NewUserName(name string) (UserName, error) {
	length := utf8.RuneCountInString(name)
	if length <= 2 || length >= 16 {
		return "", oops.BadRequest{Message: "invalid number of characters"}
	}

	return UserName(name), nil
}

func (vo UserName) String() string {
	return string(vo)
}

type Password string

func NewPassword(password string) (Password, error) {
	length := utf8.RuneCountInString(password)
	if length <= 8 || length >= 16 {
		return "", oops.BadRequest{Message: "invalid number of characters"}
	}

	return Password(password), nil
}

func (vo Password) String() string {
	return string(vo)
}

func (vo Password) IsSame(from string) bool {
	return string(vo) == from
}
