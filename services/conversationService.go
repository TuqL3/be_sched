package services

import (
	"server/dtos/conversation"
	"server/interface/Repository"
	"server/interface/Service"
	"server/models"
)

type ConversationService struct {
	ConversationService Repository.ConversationRepositoryInterface
}

func (c *ConversationService) FindOrCreateConversation(dto *conversation.FindOrCreateDto) (*models.Conversation, error) {
	return c.ConversationService.FindOrCreateConversation(dto)
}

func (c *ConversationService) UpdateLastMessageConversation(dto *conversation.UpdateLastMessageDto, conversationId uint) (*models.Conversation, error) {
	return c.ConversationService.UpdateLastMessageConversation(dto, conversationId)
}

func NewConversationService(conversation Repository.ConversationRepositoryInterface) Service.ConversationServiceInterface {
	return &ConversationService{
		ConversationService: conversation,
	}
}
