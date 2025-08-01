package response

const (
	ErrCodeSuccess      = 20001 //Success
	ErrCodeParamInvalid = 20003 //Param invalid
	ErrInvalidToken     = 20004 //Invalid token
	// err for Register
	ErrCodeRegisterEmailInvalid = 50002 //Email is invalid
	ErrCodeRegisterEmailExist   = 50001 //Email already exists
)

var msg = map[int]string{
	ErrCodeSuccess:              "Success",
	ErrCodeParamInvalid:         "Email is valid",
	ErrInvalidToken:             "Invalid token",
	ErrCodeRegisterEmailInvalid: "Email is invalid",
	ErrCodeRegisterEmailExist:   "Email already exists",
}
