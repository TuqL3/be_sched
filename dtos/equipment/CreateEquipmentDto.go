package equipment

import "github.com/go-playground/validator/v10"

type CreateEquipmentDto struct {
	Name            string `json:"name" gorm:"not null"`
	EquipmentTypeID uint   `json:"equipment_type_id" gorm:"not null"`
	RoomID          uint   `json:"room_id" gorm:"not null"`
	Status          string `json:"status" gorm:"not null"`
}

func (c *CreateEquipmentDto) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
