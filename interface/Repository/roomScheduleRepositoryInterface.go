package Repository

import (
	"server/dtos/schedule"
	"server/models"
	"server/utils"
)

type RoomScheduleRepositoryInterface interface {
	CreateSchedule(createScheduleDto *schedule.CreateRoomScheduleDto) (*models.Schedule, error)
	UpdateSchedule(roomScheduleId uint, dto schedule.UpdateRoomSchedule) (*models.Schedule, error)
	DeleteSchedule(roomScheduleId uint) error
	GetAllSchedules() ([]*models.Schedule, error)
	GetCountScheduleRoom() ([]*utils.ScheduleRoomCount, error)
}
