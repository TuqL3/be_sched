package models

import (
	"gorm.io/gorm"
	"time"
)

type ReportStatus string

const (
	ReportPending  ReportStatus = "pending"
	ReportResolved ReportStatus = "resolved"
	ReportRejected ReportStatus = "rejected"
)

type Report struct {
	ID          uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	RoomID      uint            `json:"room_id" gorm:"not null"`
	UserID      uint            `json:"user_id" gorm:"not null"`
	EquipmentID uint            `json:"equipment_id" gorm:"not null"`
	Content     string          `json:"content" gorm:"not null"`
	Status      ReportStatus    `json:"status" gorm:"not null"`
	Room        Room            `json:"room" gorm:"foreignKey:RoomID"`
	User        User            `json:"user" gorm:"foreignKey:UserID"`
	Equipment   Equipment       `json:"equipment" gorm:"foreignKey:EquipmentID"`
	CreatedAt   time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (*Report) TableName() string {
	return "report"
}
