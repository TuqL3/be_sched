package models

import (
	"gorm.io/gorm"
	"time"
)

type ScheduleStatus string

const (
	ScheduleActive   ScheduleStatus = "active"
	ScheduleInactive ScheduleStatus = "inactive"
)

type Schedule struct {
	ID          uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	RoomID      uint            `json:"room_id" gorm:"not null"`
	UserID      uint            `json:"user_id" gorm:"not null"`
	StartTime   time.Time       `json:"start_time" gorm:"not null"`
	EndTime     time.Time       `json:"end_time" gorm:"not null"`
	Status      ScheduleStatus  `json:"status" gorm:"not null"`
	Room        Room            `json:"room" gorm:"foreignKey:RoomID"`
	User        User            `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Description string          `json:"description" gorm:"not null"`
	Title       string          `json:"title"`
}

func (*Schedule) TableName() string {
	return "schedule"
}
