package repository

import (
	"time"

	"github.com/MSHR-Dec/task/go_backend/internal/domain/entity"
	"github.com/MSHR-Dec/task/go_backend/internal/domain/vo"
)

type UserRepository interface {
	Save(user entity.User, now time.Time) error
	FindByName(name vo.UserName) (entity.User, error)
}
