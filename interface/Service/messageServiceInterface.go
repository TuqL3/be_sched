package Service

import (
	"server/dtos/message"
	"server/models"
)

type MessageServiceInterface interface {
	SendMessage(dto *message.SendMessageDTO) (*models.Message, error)
	GetMessageByConversationId(conversationId uint) ([]models.Message, error)
}
