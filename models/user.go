package models

import (
	"gorm.io/gorm"
	"time"
)

type Role string

const (
	Admin     Role = "admin"
	GiangVien Role = "giang_vien"
	TrucBan   Role = "truc_ban"
	GiamDoc   Role = "giam_doc"
)

type User struct {
	ID        uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string          `json:"username" gorm:"unique;not null" validate:"required,min=3,max=50"`
	Password  string          `json:"password" gorm:"not null" validate:"required,min=8"`
	FullName  string          `json:"full_name" gorm:"not null" validate:"required"`
	Email     string          `json:"email" gorm:"unique;not null" validate:"required,email"`
	Phone     string          `json:"phone" gorm:"not null" validate:"required"`
	Role      Role            `json:"role" gorm:"not null" validate:"role"`
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (*User) TableName() string {
	return "users"
}
