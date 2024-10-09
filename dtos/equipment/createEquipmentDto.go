package equipment

import "github.com/go-playground/validator/v10"

type CreateEquipmentDto struct {
	EquipmentName string `json:"equipment_name" gorm:"not null"`
	RoomID        uint   `json:"room_id" gorm:"not null"`
	Status        string `json:"status" gorm:"not null"`
	Qty           int    `json:"qty" gorm:"not null"`
}

func (u *CreateEquipmentDto) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
