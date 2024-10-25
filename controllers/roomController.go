package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/dtos/room"
	"server/interface/Service"
	"server/utils"
	"strconv"
)

type RoomController struct {
	roomService Service.RoomServiceInterface
}

func NewRoomController(roomService Service.RoomServiceInterface) *RoomController {
	return &RoomController{
		roomService: roomService,
	}
}

func (r *RoomController) CreateRoom(c *gin.Context) {
	var roomCreateDto room.CreateRoomDto
	if err := c.ShouldBind(&roomCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := roomCreateDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	room, err := r.roomService.CreateRoom(&roomCreateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Create room failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, &utils.Response{
		Status:  http.StatusCreated,
		Message: "Create room successfully",
		Data:    room,
		Error:   "",
	})
}

func (r *RoomController) DeleteRoom(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	if err := r.roomService.DeleteRoom(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Delete room failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Delete room successfully",
		Data:    nil,
		Error:   "",
	})
	return
}

func (r *RoomController) UpdateRoom(c *gin.Context) {
	var roomUpdateDto room.UpdateRoomDto
	roomId, err := strconv.Atoi(c.Param("roomId"))

	if err := c.ShouldBind(&roomUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	room, err := r.roomService.UpdateRoom(uint(roomId), roomUpdateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Update room failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Update room successfully",
		Data:    room,
		Error:   "",
	})
}

func (r *RoomController) GetAllRoom(c *gin.Context) {
	rooms, err := r.roomService.GetAllRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Get all rooms failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Get all rooms successfully",
		Data:    rooms,
		Error:   "",
	})
	return
}

func (r *RoomController) GetRoomById(c *gin.Context) {
	roomId, err := strconv.ParseInt(c.Param("roomId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	room, err := r.roomService.GetRoomById(uint(roomId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Room get error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Room get successfully",
		Data:    room,
		Error:   "",
	})

}
