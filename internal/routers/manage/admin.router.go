package manage

import (
	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	//public router
	adminRouterPublic := Router.Group("/admin/user")
	{
		adminRouterPublic.POST("/login")
	}
	// private router
	adminRouterPrivate := Router.Group("/admin/user")
	// userAdminRouterPrivate.Use(middlewares.AuthMiddleware())
	// userAdminRouterPrivate.Use(Authen())
	// userAdminRouterPrivate.Use(Permission())

	{
		adminRouterPrivate.GET("/active_user")
	}

}
