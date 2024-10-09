package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/dtos/roomSchedule"
	"server/interface/Service"
	"server/utils"
	"strconv"
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
	var roomScheduleDto roomSchedule.CreateRoomScheduleDto
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

	roomSchedule, err := r.roomScheduleService.CreateRoomSchedule(&roomScheduleDto)
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

	if err := r.roomScheduleService.DeleteRoomSchedule(id); err != nil {
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
	var roomScheduleDto roomSchedule.UpdateRoomSchedule
	roomScheduleId, err := strconv.Atoi(c.Param("roomScheduleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid id",
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
	if err := roomScheduleDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	roomSchedule, err := r.roomScheduleService.UpdateRoomSchedule(roomScheduleId, roomScheduleDto)
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
		Message: "Updated",
		Data:    roomSchedule,
		Error:   "",
	})

}

func (r *RoomScheduleController) GetAllRoomSchedule(c *gin.Context) {
	roomSchedule, err := r.roomScheduleService.GetAllRoomSchedules()
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
