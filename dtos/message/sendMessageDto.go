package message

import "github.com/go-playground/validator/v10"

type SendMessageDTO struct {
	ConversationID uint   `gorm:"not null" json:"conversation_id"`
	SenderID       uint   `gorm:"not null" json:"sender_id"`
	ReceiverID     uint   `gorm:"not null" json:"receiver_id"`
	Content        string `gorm:"type:text;not null" json:"content"`
}

func (c *SendMessageDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
