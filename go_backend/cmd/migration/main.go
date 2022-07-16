package main

import (
	"github.com/MSHR-Dec/task/go_backend/internal/infrastructure"
	. "github.com/MSHR-Dec/task/go_backend/internal/interface/gormrepository/model"
)

func main() {
	gdb := infrastructure.NewMySQLConnection(infrastructure.Environment)

	gdb.AutoMigrate(
		&User{},
	)
}
