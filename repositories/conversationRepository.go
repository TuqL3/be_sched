package repositories

import (
	"fmt"
	"server/dtos/conversation"
	"server/interface/Repository"
	"server/models"

	"gorm.io/gorm"
)

type ConversationRepository struct {
	DB *gorm.DB
}

func (r *ConversationRepository) GetAllConversation() ([]*models.Conversation, error) {
	var conversations []*models.Conversation
	if err := r.DB.Preload("Messages").Preload("Sender").Preload("Receiver").Find(&conversations).Error; err != nil {
		return nil, err
	}
	return conversations, nil
}

func (r *ConversationRepository) GetConversationById(conversationId uint) (*models.Conversation, error) {
	var conversation models.Conversation
	if err := r.DB.Table("conversation").Where("id = ?", conversationId).Preload("Messages").First(&conversation).Error; err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *ConversationRepository) FindOrCreateConversation(dto *conversation.FindOrCreateDto) (*models.Conversation, error) {
	var conversation models.Conversation

	err := r.DB.Table("conversation").Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
		dto.SenderID, dto.ReceiverID, dto.ReceiverID, dto.SenderID).Preload("Messages").First(&conversation).Error

	if err == gorm.ErrRecordNotFound {
		var senderExists bool
		if err := r.DB.Table("user").Select("count(*) > 0").Where("id = ?", dto.SenderID).Find(&senderExists).Error; err != nil {
			return nil, err
		}
		if !senderExists {
			return nil, fmt.Errorf("SenderID %d không tồn tại", dto.SenderID)
		}

		var receiverExists bool
		if err := r.DB.Table("user").Select("count(*) > 0").Where("id = ?", dto.ReceiverID).Find(&receiverExists).Error; err != nil {
			return nil, err
		}
		if !receiverExists {
			return nil, fmt.Errorf("ReceiverID %d không tồn tại", dto.ReceiverID)
		}

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
