package infrastructure

import (
	"github.com/MSHR-Dec/task/go_backend/internal/domain/factory"
	"github.com/MSHR-Dec/task/go_backend/internal/domain/service"
	"github.com/MSHR-Dec/task/go_backend/internal/interface/controller"
	"github.com/MSHR-Dec/task/go_backend/internal/interface/gormrepository"
	"github.com/MSHR-Dec/task/go_backend/internal/usecase/interactor"
)

func injectUser() controller.UserController {
	conn := NewMySQLConnection(Environment)
	userFactory := factory.NewUserFactory()
	userRepo := gormrepository.NewUserRepository(conn)
	userService := service.NewUserService(userRepo)
	userInteractor := interactor.NewUserInteractor(userFactory, userRepo, userService)

	return controller.NewUserController(userInteractor)
}
