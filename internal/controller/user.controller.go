package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quanminhvu/go-ecommerce-backend-api/internal/service"
	"github.com/quanminhvu/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// / uc = user controller
// Controller -> Service -> Repo -> Models -> Dbs
func (uc *UserController) GetUserByID(c *gin.Context) {
	// response.ErrorResponse(c, response.ErrCodeParamInvalid, "No needs")
	response.SuccessResponse(c, response.ErrCodeSuccess, uc.userService.GetInfoUser())
}
