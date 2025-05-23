package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	message2 "server/dtos/message"
	"server/interface/Service"
	"server/models"
	"server/utils"
	"strconv"
)

type MessageController struct {
	messageService Service.MessageServiceInterface
}

func NewMessageController(messageService Service.MessageServiceInterface) *MessageController {
	return &MessageController{
		messageService: messageService,
	}
}

func (m *MessageController) SendMessage(c *gin.Context) {
	var payload message2.SendMessageDTO

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	message, err := m.messageService.SendMessage(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	messageDTO := convertMessageToDTO(message)

	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Message sent successfully",
		Data:    messageDTO,
		Error:   "",
	})
}

func convertMessageToDTO(message *models.Message) *message2.SendMessageDTO {
	return &message2.SendMessageDTO{
		ConversationID: message.ConversationID,
		SenderID:       message.SenderID,
		ReceiverID:     message.ReceiverID,
		Content:        message.Content,
	}
}

func (r *MessageController) GetMessageByConversationId(c *gin.Context) {
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
	conversation, err := r.messageService.GetMessageByConversationId(uint(conversationId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Report get error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Report get successfully",
		Data:    conversation,
		Error:   "",
	})
}
