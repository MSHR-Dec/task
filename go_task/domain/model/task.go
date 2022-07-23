package model

import (
	"database/sql"
	"database/sql/driver"
	"time"
	"unicode/utf8"

	"github.com/MSHR-Dec/MSHR-Doc/mypkg/oops"
)

type Task struct {
	ID         uint       `gorm:"primaryKey"`
	Name       TaskName   `gorm:"embedded;not null"`
	Status     TaskStatus `gorm:"embedded;tinyint(1);not null"`
	StartAt    ScheduleAt `gorm:"not null"`
	EndAt      ScheduleAt `gorm:"not null"`
	UserID     uint
	CreatedAt  time.Time    `gorm:"not null"`
	ModifiedAt sql.NullTime `gorm:"default:NULL"`
}

func NewTask(name string,
	startAt string,
	endAt string,
	userID uint,
	now time.Time,
) (Task, error) {
	taskName, err := NewTaskName(name)
	if err != nil {
		return Task{}, err
	}

	return Task{
		Name:       taskName,
		Status:     Todo,
		StartAt:    NewScheduleAt(startAt),
		EndAt:      NewScheduleAt(endAt),
		UserID:     userID,
		CreatedAt:  now,
		ModifiedAt: sql.NullTime{},
	}, nil
}

type TaskName string

func NewTaskName(name string) (TaskName, error) {
	length := utf8.RuneCountInString(name)
	if length <= 2 || length >= 50 {
		return "", oops.BadRequest{Message: "invalid number of characters"}
	}

	return TaskName(name), nil
}

func (vo TaskName) String() string {
	return string(vo)
}

type TaskStatus int

const (
	Todo TaskStatus = iota
	Doing
	Done
)

var (
	taskStatuses = []string{"Todo", "Doing", "Done"}
)

func NewTaskStatus(status string) TaskStatus {
	for i, v := range taskStatuses {
		if status == v {
			return TaskStatus(i)
		}
	}

	return Todo
}

func (vo TaskStatus) String() string {
	return taskStatuses[vo]
}

type ScheduleAt time.Time

func NewScheduleAt(schedule string) ScheduleAt {
	const layout = "2006-01-02"
	t, _ := time.Parse(layout, schedule)
	return ScheduleAt(t)
}

func (vo ScheduleAt) Time() time.Time {
	return time.Time(vo)
}

func (vo *ScheduleAt) Scan(value interface{}) error {
	*vo = ScheduleAt(value.(time.Time))
	return nil
}

func (vo ScheduleAt) Value() (driver.Value, error) {
	return time.Time(vo), nil
}
