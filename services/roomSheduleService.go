package services

import (
	"server/dtos/schedule"
	"server/interface/Repository"
	"server/interface/Service"
	"server/models"
	"server/utils"
)

type ScheduleService struct {
	scheduleRepo Repository.RoomScheduleRepositoryInterface
}

func (r *ScheduleService) GetScheduleById(scheduleId uint) (*models.Schedule, error) {
	return r.scheduleRepo.GetScheduleById(scheduleId)
}

func (r *ScheduleService) GetcountScheduleUser() ([]*utils.ScheduleUserCount, error) {
	return r.scheduleRepo.GetcountScheduleUser()
}

func (r *ScheduleService) GetCountScheduleRoom() ([]*utils.ScheduleRoomCount, error) {
	return r.scheduleRepo.GetCountScheduleRoom()
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

func (r *ScheduleService) GetAllSchedules(roomId uint, userId uint) ([]*models.Schedule, error) {
	return r.scheduleRepo.GetAllSchedules(roomId, userId)
}

func NewSheduleService(sheduleRepository Repository.RoomScheduleRepositoryInterface) Service.RoomScheduleServiceInterface {
	return &ScheduleService{
		scheduleRepo: sheduleRepository,
	}
}
