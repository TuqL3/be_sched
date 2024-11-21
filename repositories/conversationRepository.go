package repositories

import (
	"server/dtos/conversation"
	"server/interface/Repository"
	"server/models"

	"gorm.io/gorm"
)

type ConversationRepository struct {
	DB *gorm.DB
}

func (r *ConversationRepository) FindOrCreateConversation(dto *conversation.FindOrCreateDto) (*models.Conversation, error) {
	var conversation models.Conversation

	err := r.DB.Table("conversation").Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
		dto.SenderID, dto.ReceiverID, dto.ReceiverID, dto.SenderID).First(&conversation).Error

	if err == gorm.ErrRecordNotFound {
		conversation = models.Conversation{
			SenderID:   dto.SenderID,
			ReceiverID: dto.ReceiverID,
		}
		if createErr := r.DB.Table("conversation").Create(&conversation).Error; createErr != nil {
			return nil, createErr
		}
		return &conversation, nil
	}

	return &conversation, err
}

func (r *ConversationRepository) UpdateLastMessageConversation(dto *conversation.UpdateLastMessageDto, conversationId uint) (*models.Conversation, error) {
	var conversation models.Conversation
	if err := r.DB.Table("conversation").Where("id = ?", conversationId).First(&conversation).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"last_message": dto.LastMessage,
	}
	if err := r.DB.Table("conversation").Where("id = ?", conversationId).Model(&conversation).Updates(updates).Error; err != nil {
		return nil, err
	}
	return &conversation, nil
}

func NewConversationRepository(db *gorm.DB) Repository.ConversationRepositoryInterface {
	return &ConversationRepository{
		DB: db,
	}
}
