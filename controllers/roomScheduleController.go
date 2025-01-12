package controllers

import (
	"fmt"
	"net/http"
	"server/dtos/schedule"
	"server/interface/Service"
	"server/utils"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type RoomScheduleController struct {
	roomScheduleService Service.RoomScheduleServiceInterface
}

func NewRoomScheduleController(roomScheduleService Service.RoomScheduleServiceInterface) *RoomScheduleController {
	return &RoomScheduleController{
		roomScheduleService: roomScheduleService,
	}
}

func (r *RoomScheduleController) CreateRoomSchedule(c *gin.Context) {
	var roomScheduleDto schedule.CreateRoomScheduleDto

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, &utils.Response{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
			Data:    nil,
			Error:   "User not found in context",
		})
		return
	}

	claimsMap, ok := user.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Server Error",
			Data:    nil,
			Error:   "Unable to parse user claims",
		})
		return
	}

	claims := utils.Claims{}

	if idFloat, ok := claimsMap["id"].(float64); ok {
		claims.ID = fmt.Sprintf("%.0f", idFloat)
	} else if idStr, ok := claimsMap["id"].(string); ok {
		claims.ID = idStr
	}

	if roleStr, ok := claimsMap["role"].(string); ok {
		claims.Role = roleStr
	}
	
	if err := c.ShouldBindJSON(&roomScheduleDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := roomScheduleDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	roomSchedule, err := r.roomScheduleService.CreateSchedule(&roomScheduleDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Server Error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, &utils.Response{
		Status:  http.StatusCreated,
		Message: "Created",
		Data:    roomSchedule,
		Error:   "",
	})
}

func (r *RoomScheduleController) DeleteRoomSchedule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("roomScheduleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid id",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := r.roomScheduleService.DeleteSchedule(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Server Error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Deleted",
		Data:    nil,
		Error:   "",
	})

}

func (r *RoomScheduleController) UpdateRoomSchedule(c *gin.Context) {
	var roomScheduleDto schedule.UpdateRoomSchedule

	roomScheduleId, err := strconv.Atoi(c.Param("roomScheduleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid roomScheduleId",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&roomScheduleDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if roomScheduleDto.StartTime.After(roomScheduleDto.EndTime) {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "start_time cannot be after end_time",
			Data:    nil,
			Error:   "Invalid time range",
		})
		return
	}

	if err := roomScheduleDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Validation failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	roomSchedule, err := r.roomScheduleService.UpdateSchedule(uint(roomScheduleId), roomScheduleDto)
	if err != nil {

		c.JSON(http.StatusNotFound, &utils.Response{
			Status:  http.StatusNotFound,
			Message: "Room schedule not found",
			Data:    nil,
			Error:   err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Room schedule updated successfully",
		Data:    roomSchedule,
		Error:   "",
	})
}

func (r *RoomScheduleController) GetCountScheduleRoom(c *gin.Context) {
	count, err := r.roomScheduleService.GetCountScheduleRoom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Server Error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Count",
		Data:    count,
		Error:   "",
	})
}

func (r *RoomScheduleController) GetcountScheduleUser(c *gin.Context) {
	count, err := r.roomScheduleService.GetcountScheduleUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Server Error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "count",
		Data:    count,
		Error:   "",
	})
}

func (r *RoomScheduleController) GetAllRoomSchedule(c *gin.Context) {

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	claims := user.(jwt.MapClaims)

	rawRoles, ok := claims["roles"].([]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse roles"})
		return
	}

	roles := make([]string, len(rawRoles))
	for i, role := range rawRoles {
		roles[i] = role.(string)
	}

	roomIdStr := c.Query("roomId")
	userIdStr := c.Query("userId")

	roomId, _ := strconv.ParseUint(roomIdStr, 10, 64)
	userId, _ := strconv.ParseUint(userIdStr, 10, 64)

	roomSchedule, err := r.roomScheduleService.GetAllSchedules(uint(roomId), uint(userId), roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Server Error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "GetAllRoomSchedule",
		Data:    roomSchedule,
		Error:   "",
	})
}

func (r *RoomScheduleController) GetScheduleById(c *gin.Context) {
	scheduleId, err := strconv.ParseInt(c.Param("scheduleId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	schedule, err := r.roomScheduleService.GetScheduleById(uint(uint(scheduleId)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Schedule get error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Schedule get successfully",
		Data:    schedule,
		Error:   "",
	})
}

func (r *RoomScheduleController) ImportScheduleFromExcel(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}

	// Save the file temporarily
	filePath := "./" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Call usecase to process the file
	if err := r.roomScheduleService.ImportScheduleFromExcel(filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Import successful"})
}
