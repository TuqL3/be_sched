package models

import (
	"gorm.io/gorm"
	"server/utils"
	"time"
)

type Report struct {
	ID          uint               `json:"id" gorm:"primaryKey;autoIncrement"`
	RoomID      uint               `json:"room_id" gorm:"not null"`
	UserID      uint               `json:"user_id" gorm:"not null"`
	EquipmentID uint               `json:"equipment_id" gorm:"not null"` // Chỉ định thiết bị báo cáo
	Content     string             `json:"content" gorm:"not null"`
	Status      utils.ReportStatus `json:"status" gorm:"not null"` // Trạng thái của báo cáo
	Room        Room               `json:"room" gorm:"foreignKey:RoomID"`
	User        User               `json:"user" gorm:"foreignKey:UserID"`
	Equipment   Equipment          `json:"equipment" gorm:"foreignKey:EquipmentID"` // Quan hệ với Equipment
	CreatedAt   time.Time          `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time          `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *gorm.DeletedAt    `json:"deleted_at" gorm:"index"`
}

func (*Report) TableName() string {
	return "report"
}
