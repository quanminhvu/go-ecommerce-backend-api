package manage

import (
	"github.com/gin-gonic/gin"
)

type UserAdminRouter struct {
}

func (pr *UserAdminRouter) InitUserAdminRouter(Router *gin.RouterGroup) {
	// private router
	userAdminRouterPrivate := Router.Group("/admin/user")
	// userAdminRouterPrivate.Use(middlewares.AuthMiddleware())
	// userAdminRouterPrivate.Use(Authen())
	// userAdminRouterPrivate.Use(Permission())

	{
		userAdminRouterPrivate.POST("/active_user")
	}

}
