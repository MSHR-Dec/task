package repository

import (
	"github.com/MSHR-Dec/task/go_task/domain/model"
)

type TaskRepository interface {
	Save(task model.Task) (int, error)
	FindByID(id uint) (model.Task, error)
	ListByUserID(userID uint) ([]model.Task, error)
	Update(task model.Task) error
}
