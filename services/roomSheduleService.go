package services

import (
	"server/dtos/schedule"
	"server/interface/Repository"
	"server/interface/Service"
	"server/models"
)

type RoomSheduleService struct {
	roomScheduleRepo Repository.RoomScheduleRepositoryInterface
}

func (r *RoomSheduleService) CreateRoomSchedule(createRoomScheduleDto *schedule.CreateRoomScheduleDto) (*models.RoomSchedule, error) {
	return r.roomScheduleRepo.CreateRoomSchedule(createRoomScheduleDto)
}

func (r *RoomSheduleService) UpdateRoomSchedule(roomScheduleId uint, dto schedule.UpdateRoomSchedule) (*models.RoomSchedule, error) {
	return r.roomScheduleRepo.UpdateRoomSchedule(roomScheduleId, dto)
}

func (r *RoomSheduleService) DeleteRoomSchedule(roomScheduleId uint) error {
	return r.roomScheduleRepo.DeleteRoomSchedule(roomScheduleId)
}

func (r *RoomSheduleService) GetAllRoomSchedules() ([]*models.RoomSchedule, error) {
	return r.roomScheduleRepo.GetAllRoomSchedules()
}

func NewRoomSheduleService(roomScheduleRepository Repository.RoomScheduleRepositoryInterface) Service.RoomScheduleServiceInterface {
	return &RoomSheduleService{
		roomScheduleRepo: roomScheduleRepository,
	}
}
