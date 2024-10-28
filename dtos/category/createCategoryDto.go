package category

import "github.com/go-playground/validator/v10"

type CreateCategoryDto struct {
	Name string `json:"name" gorm:"not null"`
}

func (u *CreateCategoryDto) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
