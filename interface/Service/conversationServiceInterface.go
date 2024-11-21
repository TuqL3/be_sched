package Service

import (
	"server/dtos/conversation"
	"server/models"
)

type ConversationServiceInterface interface {
	FindOrCreateConversation(dto *conversation.FindOrCreateDto) (*models.Conversation, error)
	UpdateLastMessageConversation(dto *conversation.UpdateLastMessageDto, conversationId uint) (*models.Conversation, error)
}
