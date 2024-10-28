package user

import (
	"github.com/go-playground/validator/v10"
	"server/models"
)

type UserRegister struct {
	Username string        `json:"username" gorm:"unique;not null" validate:"required,min=3,max=50"`
	Password string        `json:"password" gorm:"not null" validate:"required,min=8"`
	FullName string        `json:"full_name" gorm:"not null" validate:"required"`
	Email    string        `json:"email" gorm:"unique;not null" validate:"required,email"`
	Phone    string        `json:"phone" gorm:"not null" validate:"required"`
	Roles    []models.Role `json:"roles" gorm:"many2many:user_roles"`
}

func (u *UserRegister) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
