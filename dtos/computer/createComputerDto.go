package computer

import "github.com/go-playground/validator/v10"

type CreateComputerDto struct {
	Name       string `json:"name" gorm:"not null"`
	RoomID     uint   `json:"room_id" gorm:"not null"`
	CategoryID uint   `json:"category_id" gorm:"not null"`
	Status     string `json:"status" gorm:"not null"`
}

func (u *CreateComputerDto) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
