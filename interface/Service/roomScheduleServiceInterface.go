package Service

import (
	"server/dtos/schedule"
	"server/models"
	"server/utils"
)

type RoomScheduleServiceInterface interface {
	CreateSchedule(createScheduleDto *schedule.CreateRoomScheduleDto) (*models.Schedule, error)
	UpdateSchedule(roomScheduleId uint, dto schedule.UpdateRoomSchedule) (*models.Schedule, error)
	DeleteSchedule(roomScheduleId uint) error
	GetAllSchedules(roomId uint, userId uint, roles []string) ([]*models.Schedule, error)
	GetCountScheduleRoom() ([]*utils.ScheduleRoomCount, error)
	GetcountScheduleUser() ([]*utils.ScheduleUserCount, error)
	GetScheduleById(scheduleId uint) (*models.Schedule, error)
}
