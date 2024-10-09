package report

import "github.com/go-playground/validator/v10"

type CreateReportDto struct {
	UserID      uint   `json:"user_id" gorm:"not null"`
	RoomID      uint   `json:"room_id" gorm:"not null"`
	EquipmentID uint   `json:"equipment_id" gorm:"default:null"`
	Description string `json:"description" gorm:"type:text;not null"`
	Status      string `json:"status" gorm:"not null"`
}

func (c *CreateReportDto) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
