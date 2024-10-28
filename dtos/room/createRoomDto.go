package room

import "github.com/go-playground/validator/v10"

type CreateRoomDto struct {
	Name   string `json:"name" gorm:"not null"`
	Status string `json:"status" gorm:"not null"`
}

func (u *CreateRoomDto) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
