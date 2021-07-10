package model

// UserDTO validate 기능은 github.com/go-playground/validator 참고
type UserDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewUserDTO() *UserDTO {
	return &UserDTO{}
}
