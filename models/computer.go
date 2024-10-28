package models

import (
	"server/utils"
	"time"

	"gorm.io/gorm"
)

type Computer struct {
	ID         uint                  `json:"id" gorm:"primaryKey;autoIncrement"`
	RoomID     uint                  `json:"room_id" gorm:"not null"`
	Room       Room                  `json:"room"`
	CategoryID uint                  `json:"category_id"`
	Category   Category              `json:"category"`
	NameType   string                `json:"name_type" gorm:"default:'computer'"`
	Status     utils.EquipmentStatus `json:"status" gorm:"type:equipment_status;not null"`
	CreatedAt  time.Time             `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time             `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt        `json:"deleted_at" gorm:"index"`
	Name       string                `json:"name" gorm:"not null"`
}

func (c *Computer) TableName() string {
	return "computer"
}
