package roomSchedule

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type CreateRoomScheduleDto struct {
	RoomID    uint      `json:"room_id" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	StartTime time.Time `json:"start_time" gorm:"not null"`
	EndTime   time.Time `json:"end_time" gorm:"not null"`
	Status    string    `json:"status" gorm:"not null"`
	Title     string    `json:"title" gorm:"not null"`
}

func (u *CreateRoomScheduleDto) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
