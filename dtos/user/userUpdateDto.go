package user

import (
	"github.com/go-playground/validator/v10"
	"mime/multipart"
)

type UpdateUserDto struct {
	FullName  string                `form:"full_name" json:"full_name" binding:"required"`
	Email     string                `form:"email" json:"email" binding:"required,email"`
	Phone     string                `form:"phone" json:"phone" binding:"required"`
	ImageFile *multipart.FileHeader `form:"image"`
	Roles     []uint                `form:"roles"`
	Bio       string                `form:"bio" json:"bio"`
	Github    string                `form:"github" json:"github"`
	Facebook  string                `form:"facebook" json:"facebook"`
	Instagram string                `form:"instagram" json:"instagram"`
}

func (u *UpdateUserDto) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
