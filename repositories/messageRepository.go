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

func (m MessageRepository) GetMessageByConversationId(conversationId uint) ([]models.Message, error) {
	var messages []models.Message
	if err := m.DB.Table("message").
		Where("conversation_id = ?", conversationId).
		Preload("Sender").   // Tải thông tin người gửi
		Preload("Receiver"). // Tải thông tin người nhận
		Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (m MessageRepository) SendMessage(dto *message.SendMessageDTO) (*models.Message, error) {
	if err := m.DB.Table("message").Create(dto).Error; err != nil {
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
