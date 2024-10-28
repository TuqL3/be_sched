package role

import "github.com/go-playground/validator/v10"

type CreateRoleDto struct {
	RoleName string `json:"role_name" gorm:"unique;not null"`
}

func (c *CreateRoleDto) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
