package gormrepository

import (
	"errors"

	"github.com/MSHR-Dec/MSHR-Doc/mypkg/oops"
	"gorm.io/gorm"

	"github.com/MSHR-Dec/task/go_task/domain/model"
)

type TaskRepository struct {
	gdb *gorm.DB
}

func NewTaskRepository(gdb *gorm.DB) TaskRepository {
	return TaskRepository{
		gdb: gdb,
	}
}

func (r TaskRepository) Save(task model.Task) (int, error) {
	tx := r.gdb.Begin()

	if err := tx.Save(&task).Error; err != nil {
		tx.Rollback()
		return 0, oops.InternalServerError{Message: err.Error()}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return 0, oops.InternalServerError{Message: err.Error()}
	}

	return int(task.ID), nil
}

func (r TaskRepository) FindByID(id uint) (model.Task, error) {
	var task model.Task
	if err := r.gdb.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Task{}, oops.NotFound{Message: err.Error()}
		} else {
			return model.Task{}, oops.InternalServerError{Message: err.Error()}
		}
	}

	return task, nil
}

func (r TaskRepository) ListByUserID(userID uint) ([]model.Task, error) {
	var tasks []model.Task
	if result := r.gdb.Where("user_id = ?", userID).Find(&tasks); result.Error != nil {
		return nil, oops.InternalServerError{Message: result.Error.Error()}
	} else if result.RowsAffected <= 0 {
		return nil, oops.NotFound{Message: gorm.ErrRecordNotFound.Error()}
	}

	return tasks, nil
}

func (r TaskRepository) Update(task model.Task) error {
	tx := r.gdb.Begin()

	if err := tx.Save(&task).Error; err != nil {
		tx.Rollback()
		return oops.InternalServerError{Message: err.Error()}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return oops.InternalServerError{Message: err.Error()}
	}

	return nil
}
