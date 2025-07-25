package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/quanminhvu/go-ecommerce-backend-api/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping/:name", c.NewPongController().Pong)
		v1.GET("/user/1", c.NewUserController().GetUserByID)
	}
	return r
}
