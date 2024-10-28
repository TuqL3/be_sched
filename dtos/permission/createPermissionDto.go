package permission

import "github.com/go-playground/validator/v10"

type CreatePermissionDto struct {
	PermissionName string `json:"permission_name" gorm:"unique;not null"`
}

func (c *CreatePermissionDto) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
