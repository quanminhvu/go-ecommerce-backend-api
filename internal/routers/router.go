package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping/:name", Pong)
		v1.GET("/ping", Pong)
	}
	return r
}
func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "quanminhvu")
	// c.ShouldBindJSON()
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...ping to " + name,
		"uid":     uid,
		"user":    []string{"Cr7", "M10", "Messi"},
	})
}
