package response

const (
	ErrCodeSuccess      = 20001 //Success
	ErrCodeParamInvalid = 20003 //Param invalid
	ErrInvalidToken     = 20004 //Invalid token
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is valid",
	ErrInvalidToken:     "Invalid token",
}
