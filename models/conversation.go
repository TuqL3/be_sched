package models

import "time"

type Conversation struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	SenderID    uint      `gorm:"not null" json:"sender_id"`
	ReceiverID  uint      `gorm:"not null" json:"receiver_id"`
	LastMessage string    `gorm:"type:text" json:"last_message"`
	UpdatedAt   time.Time `json:"updated_at"`

	Sender   User `gorm:"foreignKey:SenderID" json:"sender"`
	Receiver User `gorm:"foreignKey:ReceiverID" json:"receiver"`

	Messages []Message `gorm:"foreignKey:ConversationID" json:"messages"`
}

func (*Conversation) TableName() string {
	return "conversation"
}
