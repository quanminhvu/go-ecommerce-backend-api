package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "quanminhvu")
	// c.ShouldBindJSON()
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...ping to " + name,
		"uid":     uid,
		"user":    []string{"Cr7", "M10", "Messi"},
	})
}
