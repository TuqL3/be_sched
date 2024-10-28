package models

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID          uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	RoleName    string          `json:"role_name" gorm:"unique;not null"`
	Permissions []Permission    `json:"permissions" gorm:"many2many:role_permissions"`
	Users       []User          `json:"users" gorm:"many2many:user_roles"`
	CreatedAt   time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (*Role) TableName() string {
	return "role"
}
