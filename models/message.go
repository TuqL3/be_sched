package models

import (
	"time"
)

type Message struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	ConversationID uint      `gorm:"not null" json:"conversation_id"`
	SenderID       uint      `gorm:"not null" json:"sender_id"`
	ReceiverID     uint      `gorm:"not null" json:"receiver_id"`
	Content        string    `gorm:"type:text;not null" json:"content"`
	CreatedAt      time.Time `json:"created_at"`

	Conversation Conversation `gorm:"foreignKey:ConversationID" json:"-"`
	Sender       User         `gorm:"foreignKey:SenderID" json:"sender"`
	Receiver     User         `gorm:"foreignKey:ReceiverID" json:"receiver"`
}

func (*Message) TableName() string {
	return "message"
}
