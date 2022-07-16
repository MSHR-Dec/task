package adapter

import (
	"github.com/MSHR-Dec/task/go_task/application"
	"github.com/MSHR-Dec/task/go_task/domain/service"
	"github.com/MSHR-Dec/task/go_task/interfaces/controller"
	"github.com/MSHR-Dec/task/go_task/interfaces/gormrepository"
)

func injectUser() controller.UserController {
	conn := NewMySQLConnection(Environment)
	userRepo := gormrepository.NewUserRepository(conn)
	userService := service.NewUserService(userRepo)
	userInteractor := application.NewUserInteractor(userRepo, userService)

	return controller.NewUserController(userInteractor)
}
