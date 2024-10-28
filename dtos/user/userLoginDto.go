package user

import "github.com/go-playground/validator/v10"

type UserLoginDto struct {
	Username string `json:"username" gorm:"unique;not null" validate:"required,min=3,max=50"`
	Password string `json:"password" gorm:"not null" validate:"required,min=8"`
}

func (u *UserLoginDto) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
