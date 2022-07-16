package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID                     uint         `gorm:"primaryKey"`
	UUID                   string       `gorm:"unique;not null"`
	Name                   string       `gorm:"unique;not null"`
	Password               string       `gorm:"not null"`
	LastPasswordModifiedAt time.Time    `gorm:"not null"`
	CreatedAt              time.Time    `gorm:"not null"`
	ModifiedAt             sql.NullTime `gorm:"default:NULL"`
}
