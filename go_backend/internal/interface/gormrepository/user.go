package gormrepository

import (
	"database/sql"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/MSHR-Dec/task/go_backend/internal/domain/entity"
	"github.com/MSHR-Dec/task/go_backend/internal/domain/repository"
	"github.com/MSHR-Dec/task/go_backend/internal/domain/vo"
	"github.com/MSHR-Dec/task/go_backend/internal/interface/gormrepository/model"
	"github.com/MSHR-Dec/task/go_backend/pkg/oops"
)

type UserRepository struct {
	gdb *gorm.DB
}

func NewUserRepository(gdb *gorm.DB) repository.UserRepository {
	return UserRepository{
		gdb: gdb,
	}
}

func (r UserRepository) toModel(user entity.User, now time.Time, isUpdate bool) *model.User {
	modified := sql.NullTime{}
	if isUpdate {
		modified.Time = now
		modified.Valid = isUpdate
	}

	return &model.User{
		UUID:                   user.UUID.String(),
		Name:                   user.Name.String(),
		Password:               user.Password.String(),
		LastPasswordModifiedAt: now,
		CreatedAt:              now,
		ModifiedAt:             modified,
	}
}

func (r UserRepository) toEntity(user model.User) entity.User {
	return entity.User{
		UUID:                   vo.UUID(user.UUID),
		Name:                   vo.UserName(user.Name),
		Password:               vo.Password(user.Password),
		LastPasswordModifiedAt: user.LastPasswordModifiedAt,
	}
}

func (r UserRepository) Save(user entity.User, now time.Time) error {
	tx := r.gdb.Begin()

	if err := tx.Create(r.toModel(user, now, false)).Error; err != nil {
		tx.Rollback()
		return oops.InternalServerError{Message: err.Error()}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return oops.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (r UserRepository) FindByName(name vo.UserName) (entity.User, error) {
	var user model.User
	if err := r.gdb.Where("name = ?", name.String()).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, oops.NotFound{Message: err.Error()}
		} else {
			return entity.User{}, oops.InternalServerError{Message: err.Error()}
		}
	}

	return r.toEntity(user), nil
}
