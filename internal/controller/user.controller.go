package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quanminhvu/go-ecommerce-backend-api/internal/service"
	"github.com/quanminhvu/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(c, response.ErrCodeSuccess, result)

}

// // / uc = user controller
// // Controller -> Service -> Repo -> Models -> Dbs
// func (uc *UserController) GetUserByID(c *gin.Context) {
// 	fmt.Println("-----> My Handler")
// 	// response.ErrorResponse(c, response.ErrCodeParamInvalid, "No needs")
// 	response.SuccessResponse(c, response.ErrCodeSuccess, uc.userService.GetInfoUser())
// }
