package repository

import (
	"github.com/MSHR-Dec/task/go_task/domain/model"
)

type UserRepository interface {
	Save(model.User) error
	FindByName(name model.UserName) (model.User, error)
}
