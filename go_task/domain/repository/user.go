package repository

import (
	"github.com/MSHR-Dec/task/go_task/domain/model"
)

type UserRepository interface {
	Save(model.User) (int, error)
	FindByID(id uint) (model.User, error)
	FindByName(name model.UserName) (model.User, error)
	Update(user model.User) error
}
