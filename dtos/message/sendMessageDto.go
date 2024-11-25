package message

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type SendMessageDTO struct {
	ConversationID uint      `gorm:"not null" json:"conversation_id"`
	SenderID       uint      `gorm:"not null" json:"sender_id"`
	ReceiverID     uint      `gorm:"not null" json:"receiver_id"`
	Content        string    `gorm:"type:text;not null" json:"content"`
	CreatedAt      time.Time `json:"created_at"`
}

func (c *SendMessageDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
