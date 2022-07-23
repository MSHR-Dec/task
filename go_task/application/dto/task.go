package dto

import (
	"time"
)

type TaskAddInput struct {
	Name    string `json:"name" binding:"required"`
	StartAt string `json:"start_at" binding:"required"`
	EndAt   string `json:"end_at" binding:"required"`
	UserID  int
}

type TaskAddOutput struct {
	ID int `json:"id"`
}

type TaskEditInput struct {
	ID      int    `json:"id" binding:"required"`
	Name    string `json:"name,omitempty"`
	Status  string `json:"status,omitempty"`
	StartAt string `json:"start_at,omitempty"`
	EndAt   string `json:"end_at,omitempty"`
}

type TaskListInput struct {
	UserID int
}

type TaskListOutput struct {
	Tasks []TaskOutput `json:"tasks"`
}

type TaskOutput struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	StartAt    time.Time `json:"start_at"`
	EndAt      time.Time `json:"end_at"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}
