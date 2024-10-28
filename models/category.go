package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID            uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string         `json:"name" gorm:"not null"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Computers     []Computer     `json:"computers" gorm:"foreignKey:RoomID"`
	AirConditions []AirCondition `json:"airConditions" gorm:"foreignKey:RoomID"`
	TAndChs       []TandCh       `json:"tandChs" gorm:"foreignKey:RoomID"`
}

func (*Category) TableName() string {
	return "categories"
}
