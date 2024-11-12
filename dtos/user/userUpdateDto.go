package user

import (
	"github.com/go-playground/validator/v10"
)

type UpdateUserDto struct {
	FullName string `json:"full_name" validate:"omitempty"`
	Email    string `json:"email" validate:"omitempty,email"`
	ImageUrl string `json:"image_url"`
	Phone    string `json:"phone" validate:"omitempty"`
	Roles    []uint `json:"roles"`
}

func (u *UpdateUserDto) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
