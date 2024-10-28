package Service

import (
	"server/dtos/schedule"
	"server/models"
)

type RoomScheduleServiceInterface interface {
	CreateRoomSchedule(createRoomScheduleDto *schedule.CreateRoomScheduleDto) (*models.RoomSchedule, error)
	UpdateRoomSchedule(roomScheduleId uint, dto schedule.UpdateRoomSchedule) (*models.RoomSchedule, error)
	DeleteRoomSchedule(roomScheduleId uint) error
	GetAllRoomSchedules() ([]*models.RoomSchedule, error)
}
