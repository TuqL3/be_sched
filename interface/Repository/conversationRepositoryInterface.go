package Repository

import (
	"server/dtos/conversation"
	"server/models"
)

type ConversationRepositoryInterface interface {
	FindOrCreateConversation(dto *conversation.FindOrCreateDto) (*models.Conversation, error)
	UpdateLastMessageConversation(dto *conversation.UpdateLastMessageDto, conversationId uint) (*models.Conversation, error)
}
