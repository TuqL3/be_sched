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

func (s *ConversationService) GetAllConversation() ([]*models.Conversation, error) {
	return s.ConversationService.GetAllConversation()
}

func (s *ConversationService) GetConversationById(conversationId uint) (*models.Conversation, error) {
	return s.ConversationService.GetConversationById(conversationId)
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
