package models

import (
	"gorm.io/gorm"
	"time"
)

type EquipmentType struct {
	ID         uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string          `json:"name" gorm:"unique;not null"`
	Equipments []Equipment     `json:"equipments" gorm:"foreignKey:EquipmentTypeID"`
	CreatedAt  time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (*EquipmentType) TableName() string {
	return "equipmentType"
}
