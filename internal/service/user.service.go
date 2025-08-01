package service

import (
	"github.com/quanminhvu/go-ecommerce-backend-api/internal/repo"
	"github.com/quanminhvu/go-ecommerce-backend-api/pkg/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.userRepo.GetInfoUser()
// }

//INTERFACE VERSION

type IUserService interface {
	Register(email string, purpose string) bool
}

type userService struct {
	userRepo repo.IUserRepository
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) Register(email string, purpose string) bool {
	if us.userRepo.GetUserByEmail(email, purpose) {
		return response.ErrCodeRegisterEmailExist == response.ErrCodeRegisterEmailExist
	}
	return response.ErrCodeSuccess == response.ErrCodeSuccess
}
