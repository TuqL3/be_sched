package models

import (
	"gorm.io/gorm"
	"time"
)

type Permission struct {
	ID             uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	PermissionName string          `json:"permission_name" gorm:"unique;not null"`
	Roles          []Role          `json:"roles" gorm:"many2many:role_permissions"`
	CreatedAt      time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (*Permission) TableName() string {
	return "permission"
}
