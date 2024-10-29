package Service

import (
	"server/dtos/schedule"
	"server/models"
)

type RoomScheduleServiceInterface interface {
	CreateSchedule(createScheduleDto *schedule.CreateRoomScheduleDto) (*models.Schedule, error)
	UpdateSchedule(roomScheduleId uint, dto schedule.UpdateRoomSchedule) (*models.Schedule, error)
	DeleteSchedule(roomScheduleId uint) error
	GetAllSchedules() ([]*models.Schedule, error)
}
