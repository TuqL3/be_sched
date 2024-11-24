package services

import (
	"server/dtos/message"
	"server/interface/Repository"
	"server/interface/Service"
	"server/models"
)

type MessageService struct {
	messageRepository Repository.MessageRepositoryInterface
}

func (m MessageService) GetMessageByConversationId(conversationId uint) ([]models.Message, error) {
	return m.messageRepository.GetMessageByConversationId(conversationId)
}

func (m MessageService) SendMessage(dto *message.SendMessageDTO) (*models.Message, error) {
	return m.messageRepository.SendMessage(dto)
}

func NewMessageService(messageRepository Repository.MessageRepositoryInterface) Service.MessageServiceInterface {
	return &MessageService{
		messageRepository: messageRepository,
	}
}
