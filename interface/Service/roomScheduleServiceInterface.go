package Service

import (
	"server/dtos/roomSchedule"
	"server/models"
)

type RoomScheduleServiceInterface interface {
	CreateRoomSchedule(createRoomScheduleDto *roomSchedule.CreateRoomScheduleDto) (*models.RoomSchedule, error)
	UpdateRoomSchedule(roomScheduleId int, dto roomSchedule.UpdateRoomSchedule) (*models.RoomSchedule, error)
	DeleteRoomSchedule(roomScheduleId int) error
	GetAllRoomSchedules() ([]*models.RoomSchedule, error)
}
