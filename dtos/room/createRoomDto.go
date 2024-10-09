package room

import "github.com/go-playground/validator/v10"

type CreateRoomDto struct {
	RoomName string `json:"room_name" gorm:"not null"`
	Capacity uint   `json:"capacity" gorm:"not null"`
	Status   string `json:"status" gorm:"not null"`
}

func (u *CreateRoomDto) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
