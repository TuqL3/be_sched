package models

import (
	"gorm.io/gorm"
	"time"
)

type EquipmentStatus string

const (
	Working    EquipmentStatus = "working"
	Broken     EquipmentStatus = "broken"
	Maintained EquipmentStatus = "maintained"
)

type Equipment struct {
	ID            uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	EquipmentName string          `json:"equipment_name" gorm:"not null"`
	RoomID        uint            `json:"room_id" gorm:"not null"`
	Room          Room            `json:"room" gorm:"foreignKey:RoomID"`
	Status        EquipmentStatus `json:"status" gorm:"type:equipment_status;not null"`
	Qty           int             `json:"qty" gorm:"not null"`
	CreatedAt     time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt  `json:"deleted_at" gorm:"index"`
}

func (*Equipment) TableName() string {
	return "equipment"
}
