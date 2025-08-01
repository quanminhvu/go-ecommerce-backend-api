package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quanminhvu/go-ecommerce-backend-api/internal/wire"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	// non-dependency injection
	// ur := repo.NewUserRepository()
	// us := service.NewUserService(ur)
	// userController := controller.NewUserController(us)
	userController, _ := wire.InitUserRouteHandler()
	// use wire dependency injection

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(middlewares.AuthMiddleware())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())

	{
		userRouterPrivate.GET("/info")
	}

}
