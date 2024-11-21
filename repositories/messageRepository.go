package repositories

import (
	"gorm.io/gorm"
	"server/dtos/message"
	"server/interface/Repository"
	"server/models"
)

type MessageRepository struct {
	DB *gorm.DB
}

func (m MessageRepository) SendMessage(dto *message.SendMessageDTO) (*models.Message, error) {
	if err := m.DB.Table("messages").Create(dto).Error; err != nil {
		return nil, err
	}

	msg := &models.Message{
		SenderID:       dto.SenderID,
		ReceiverID:     dto.ReceiverID,
		Content:        dto.Content,
		ConversationID: dto.ConversationID,
	}
	return msg, nil
}

func NewMessageRepository(db *gorm.DB) Repository.MessageRepositoryInterface {
	return &MessageRepository{
		DB: db,
	}
}
