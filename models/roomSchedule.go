package models

import (
	"gorm.io/gorm"
	"time"
)

type ScheduleStatus string

const (
	Pending   ScheduleStatus = "pending"
	Approved  ScheduleStatus = "approved"
	Completed ScheduleStatus = "completed"
)

type RoomSchedule struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	RoomID    uint           `json:"room_id" gorm:"not null"`
	Room      Room           `json:"room" gorm:"foreignKey:RoomID"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	User      User           `json:"user" gorm:"foreignKey:UserID"`
	StartTime time.Time      `json:"start_time" gorm:"not null"`
	EndTime   time.Time      `json:"end_time" gorm:"not null"`
	Status    ScheduleStatus `json:"status" gorm:"type:room_schedule_status;not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Title     string         `json:"title" gorm:"title"`
}

func (*RoomSchedule) TableName() string {
	return "room_schedule"
}
