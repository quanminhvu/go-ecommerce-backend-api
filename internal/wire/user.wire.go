//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/quanminhvu/go-ecommerce-backend-api/internal/controller"
	"github.com/quanminhvu/go-ecommerce-backend-api/internal/repo"
	"github.com/quanminhvu/go-ecommerce-backend-api/internal/service"
)

func InitUserRouteHandler() (*controller.UserController, error) {
	wire.Build(
		controller.NewUserController,
		service.NewUserService,
		repo.NewUserRepository,
	)

	return new(controller.UserController), nil
}
