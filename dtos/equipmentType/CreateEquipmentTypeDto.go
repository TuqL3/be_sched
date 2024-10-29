package equipmentType

import "github.com/go-playground/validator/v10"

type CreateEquipmentTypeDto struct {
	Name string `json:"name" gorm:"unique;not null"`
}

func (c *CreateEquipmentTypeDto) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
