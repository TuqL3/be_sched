package Repository

import (
	"server/dtos/roomSchedule"
	"server/models"
)

type RoomScheduleRepositoryInterface interface {
	CreateRoomSchedule(createRoomScheduleDto *roomSchedule.CreateRoomScheduleDto) (*models.RoomSchedule, error)
	UpdateRoomSchedule(roomScheduleId uint, dto roomSchedule.UpdateRoomSchedule) (*models.RoomSchedule, error)
	DeleteRoomSchedule(roomScheduleId uint) error
	GetAllRoomSchedules() ([]*models.RoomSchedule, error)
}
