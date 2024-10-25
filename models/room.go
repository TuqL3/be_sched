package models

import (
	"gorm.io/gorm"
	"time"
)

type RoomStatus string

const (
	Available   RoomStatus = "available"
	InUse       RoomStatus = "in_use"
	Maintenance RoomStatus = "maintenance"
)

type Room struct {
	ID            uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	RoomName      string         `json:"room_name" gorm:"not null"`
	Capacity      uint           `json:"capacity" gorm:"not null"`
	Status        RoomStatus     `json:"status" gorm:"type:room_status;not null"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Computers     []Computer     `json:"computers" gorm:"foreignKey:RoomID"`
	AirConditions []AirCondition `json:"airConditions" gorm:"foreignKey:RoomID"`
	TAndChs       []TandCh       `json:"tandChs" gorm:"foreignKey:RoomID"`
}

func (*Room) TableName() string {
	return "rooms"
}
