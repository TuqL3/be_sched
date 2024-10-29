package services

import (
	"server/dtos/schedule"
	"server/interface/Repository"
	"server/interface/Service"
	"server/models"
)

type ScheduleService struct {
	scheduleRepo Repository.RoomScheduleRepositoryInterface
}

func (r *ScheduleService) CreateSchedule(createScheduleDto *schedule.CreateRoomScheduleDto) (*models.Schedule, error) {
	return r.scheduleRepo.CreateSchedule(createScheduleDto)
}

func (r *ScheduleService) UpdateSchedule(roomScheduleId uint, dto schedule.UpdateRoomSchedule) (*models.Schedule, error) {
	return r.scheduleRepo.UpdateSchedule(roomScheduleId, dto)
}

func (r *ScheduleService) DeleteSchedule(roomScheduleId uint) error {
	return r.scheduleRepo.DeleteSchedule(roomScheduleId)
}

func (r *ScheduleService) GetAllSchedules() ([]*models.Schedule, error) {
	return r.scheduleRepo.GetAllSchedules()
}

func NewSheduleService(sheduleRepository Repository.RoomScheduleRepositoryInterface) Service.RoomScheduleServiceInterface {
	return &ScheduleService{
		scheduleRepo: sheduleRepository,
	}
}
