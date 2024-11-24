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
func (r *ConversationController) GetConversaionById(c *gin.Context) {
	conversationId, err := strconv.ParseInt(c.Param("conversationId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	conversation, err := r.conversationService.GetConversationById(uint(conversationId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Conversation get error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Air condition get successfully",
		Data:    conversation,
		Error:   "",
	})

}

func (r *ConversationController) GetAllConversation(c *gin.Context) {
	equipment, err := r.conversationService.GetAllConversation()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Get equipment failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Get equipment successfully",
		Data:    equipment,
		Error:   "",
	})
	return
}
