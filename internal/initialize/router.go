package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/quanminhvu/go-ecommerce-backend-api/global"
	"github.com/quanminhvu/go-ecommerce-backend-api/internal/routers"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares
	// r.Use()

	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1")
	{
		MainGroup.GET("/ping") // tracking monitor

	}
	{
		manageRouter.InitAdminRouter(MainGroup)
		manageRouter.InitUserAdminRouter(MainGroup)

	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}

	return r
}
