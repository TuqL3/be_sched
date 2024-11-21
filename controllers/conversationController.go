package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/dtos/conversation"
	"server/interface/Service"
	"server/utils"
	"strconv"
)

type ConversationController struct {
	conversationService Service.ConversationServiceInterface
}

func NewConversationController(conversationService Service.ConversationServiceInterface) *ConversationController {
	return &ConversationController{conversationService: conversationService}
}

func (c *ConversationController) FindOrCreateConversation(ctx *gin.Context) {
	var payload struct {
		SenderID   uint `json:"sender_id"`
		ReceiverID uint `json:"receiver_id"`
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conversation, err := c.conversationService.FindOrCreateConversation(
		&conversation.FindOrCreateDto{
			SenderID:   payload.SenderID,
			ReceiverID: payload.ReceiverID,
		},
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, conversation)
}

func (c *ConversationController) UpdateLastMessage(ctx *gin.Context) {
	var conversationDto conversation.UpdateLastMessageDto
	id, err := strconv.Atoi(ctx.Param("conversationId"))
	if err := ctx.ShouldBindJSON(&conversationDto); err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	conversation, err := c.conversationService.UpdateLastMessageConversation(&conversationDto, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    conversation,
		Error:   "",
	})
}
