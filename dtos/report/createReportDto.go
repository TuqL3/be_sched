package report

import "github.com/go-playground/validator/v10"

type CreateReportDto struct {
	RoomID      uint   `json:"room_id" gorm:"not null"`
	UserID      uint   `json:"user_id" gorm:"not null"`
	EquipmentID uint   `json:"equipment_id" gorm:"not null"`
	Content     string `json:"content" gorm:"not null"`
	Status      string `json:"status" gorm:"not null"`
}

func (c *CreateReportDto) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
