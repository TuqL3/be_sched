package user

import (
	"github.com/go-playground/validator/v10"
)

type UserRegister struct {
	Username string `json:"username" binding:"required" validate:"min=3,max=50"`
	Password string `json:"password" binding:"required" validate:"min=8"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Role     string `json:"role" binding:"required" validate:"oneof=admin giang_vien truc_ban giam_doc"`
}

func (u *UserRegister) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
