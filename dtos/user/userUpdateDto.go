package user

import (
	"github.com/go-playground/validator/v10"
	"server/models"
)

type UserUpdateDto struct {
	FullName string        `json:"full_name" binding:"required"`
	Email    string        `json:"email" binding:"required,email"`
	Phone    string        `json:"phone" binding:"required"`
	Roles    []models.Role `json:"roles" gorm:"many2many:user_roles"`
}

func (u *UserUpdateDto) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
