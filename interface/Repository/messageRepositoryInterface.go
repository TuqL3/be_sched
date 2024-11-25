package Repository

import (
	"server/dtos/message"
	"server/models"
)

type MessageRepositoryInterface interface {
	SendMessage(dto *message.SendMessageDTO) (*models.Message, error)
	GetMessageByConversationId(conversationId uint) ([]models.Message, error)
}
