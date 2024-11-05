package schedule

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type CreateRoomScheduleDto struct {
	RoomID      uint      `json:"location" gorm:"not null"`
	UserID      uint      `json:"participants" gorm:"not null"`
	StartTime   time.Time `json:"start" gorm:"not null"`
	EndTime     time.Time `json:"end" gorm:"not null"`
	Status      string    `json:"status" gorm:"not null"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
}

func (u *CreateRoomScheduleDto) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
