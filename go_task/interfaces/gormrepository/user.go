package gormrepository

import (
	"errors"

	"github.com/MSHR-Dec/MSHR-Doc/mypkg/oops"
	"gorm.io/gorm"

	"github.com/MSHR-Dec/task/go_task/domain/model"
)

type UserRepository struct {
	gdb *gorm.DB
}

func NewUserRepository(gdb *gorm.DB) UserRepository {
	return UserRepository{
		gdb: gdb,
	}
}

func (r UserRepository) Save(user model.User) error {
	tx := r.gdb.Begin()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return oops.InternalServerError{Message: err.Error()}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return oops.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (r UserRepository) FindByName(name model.UserName) (model.User, error) {
	var user model.User
	if err := r.gdb.Where("name = ?", name.String()).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, oops.NotFound{Message: err.Error()}
		} else {
			return model.User{}, oops.InternalServerError{Message: err.Error()}
		}
	}

	return user, nil
}
