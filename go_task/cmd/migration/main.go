package main

import (
	"github.com/MSHR-Dec/task/go_task/adapter"
	. "github.com/MSHR-Dec/task/go_task/domain/model"
)

func main() {
	gdb := adapter.NewMySQLConnection(adapter.Environment)

	gdb.AutoMigrate(
		&Task{},
		&User{},
	)
}
