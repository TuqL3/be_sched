package models

import (
	"gorm.io/gorm"
	"time"
)

type ReportStatus string

const (
	ReportPending ReportStatus = "pending"
	InProgress    ReportStatus = "in_progress"
	Resolved      ReportStatus = "resolved"
	Rejected      ReportStatus = "rejected"
)

type Report struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      uint           `json:"user_id" gorm:"not null"`
	User        User           `json:"user" gorm:"foreignKey:UserID"`
	RoomID      uint           `json:"room_id" gorm:"not null"`
	Room        Room           `json:"room" gorm:"foreignKey:RoomID"`
	EquipmentID uint           `json:"equipment_id" gorm:"default:null"`
	Equipment   Equipment      `json:"equipment" gorm:"foreignKey:EquipmentID"`
	Description string         `json:"description" gorm:"type:text;not null"`
	Status      ReportStatus   `json:"status" gorm:"type:report_status;not null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (*Report) TableName() string {
	return "reports"
}
