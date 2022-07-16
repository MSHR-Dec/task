package factory

import (
	"time"

	"github.com/MSHR-Dec/task/go_backend/internal/domain/entity"
	"github.com/MSHR-Dec/task/go_backend/internal/domain/vo"
)

type UserFactory struct{}

func NewUserFactory() UserFactory {
	return UserFactory{}
}

func (f UserFactory) Create(name string, password string, now time.Time) (entity.User, error) {
	userName, err := vo.NewUserName(name)
	if err != nil {
		return entity.User{}, err
	}

	userPassword, err := vo.NewPassword(password)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{
		UUID:                   vo.NewUUID(),
		Name:                   userName,
		Password:               userPassword,
		LastPasswordModifiedAt: now,
	}, nil
}
