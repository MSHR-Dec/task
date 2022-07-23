package model

import (
	"database/sql"
	"time"
	"unicode/utf8"

	"github.com/MSHR-Dec/MSHR-Doc/mypkg/oops"
	"github.com/google/uuid"
)

type User struct {
	ID                     uint     `gorm:"primaryKey"`
	UUID                   UUID     `gorm:"embedded;unique;not null"`
	Name                   UserName `gorm:"embedded;unique;not null"`
	Password               Password `gorm:"embedded;not null"`
	Tasks                  []Task
	LastPasswordModifiedAt time.Time    `gorm:"not null"`
	CreatedAt              time.Time    `gorm:"not null"`
	ModifiedAt             sql.NullTime `gorm:"default:NULL"`
}

func NewUser(name string, password string, now time.Time) (User, error) {
	userName, err := NewUserName(name)
	if err != nil {
		return User{}, err
	}

	userPassword, err := NewPassword(password)
	if err != nil {
		return User{}, err
	}

	return User{
		UUID:                   NewUUID(),
		Name:                   userName,
		Password:               userPassword,
		LastPasswordModifiedAt: now,
		CreatedAt:              now,
		ModifiedAt:             sql.NullTime{},
	}, nil
}

func (e User) ShouldUpdatePassword() bool {
	return int(time.Since(e.LastPasswordModifiedAt).Hours()) > 24*7
}

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
