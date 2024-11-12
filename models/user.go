package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string          `json:"username" gorm:"unique;not null" validate:"required,min=3,max=50"`
	Password  string          `json:"password" gorm:"not null" validate:"required,min=8"`
	FullName  string          `json:"full_name" gorm:"not null" validate:"required"`
	Email     string          `json:"email" gorm:"unique;not null" validate:"required,email"`
	Phone     string          `json:"phone" gorm:"not null" validate:"required"`
	Roles     []Role          `json:"roles" gorm:"many2many:user_roles"`
	Schedules []Schedule      `json:"schedules" gorm:"foreignKey:UserID"`
	Reports   []Report        `json:"reports" gorm:"foreignKey:UserID"`
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ImageUrl  string          `json:"image_url"`
}

func (*User) TableName() string {
	return "user"
}
