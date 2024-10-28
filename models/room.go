package models

import (
	"gorm.io/gorm"
	"server/utils"
	"time"
)

type Room struct {
	ID         uint             `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string           `json:"name" gorm:"not null"`
	Status     utils.RoomStatus `json:"status" gorm:"not null"`
	Equipments []Equipment      `json:"equipments" gorm:"foreignKey:RoomID"`
	Schedules  []Schedule       `json:"schedules" gorm:"foreignKey:RoomID"`
	Reports    []Report         `json:"reports" gorm:"foreignKey:RoomID"`
	CreatedAt  time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  *gorm.DeletedAt  `json:"deleted_at" gorm:"index"`
}

func (*Room) TableName() string {
	return "room"
}
