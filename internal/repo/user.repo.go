package repo

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser() string {
// 	return "TipsGolang"
// }

//INTERFACE VERSION

type IUserRepository interface {
	GetUserByEmail(email string, purpose string) bool
}

type userRepository struct{}

// GetUserByEmail implements IUserRepository.
func (u *userRepository) GetUserByEmail(email string, purpose string) bool {
	return true // Simulating a user found by email
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
