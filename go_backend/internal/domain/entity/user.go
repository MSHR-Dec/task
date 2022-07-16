package entity

import (
	"time"

	"github.com/MSHR-Dec/task/go_backend/internal/domain/vo"
)

type User struct {
	UUID                   vo.UUID
	Name                   vo.UserName
	Password               vo.Password
	LastPasswordModifiedAt time.Time
}

func (e User) ShouldUpdatePassword() bool {
	return int(time.Since(e.LastPasswordModifiedAt).Hours()) > 24*7
}
