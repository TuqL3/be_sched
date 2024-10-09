package repositories

import (
	"errors"
	"gorm.io/gorm"
	"server/dtos/roomSchedule"
	"server/interface/Repository"
	"server/models"
	"time"
)

type RoomScheduleRepository struct {
	DB *gorm.DB
}

func (r *RoomScheduleRepository) CreateRoomSchedule(createRoomScheduleDto *roomSchedule.CreateRoomScheduleDto) (*models.RoomSchedule, error) {
	if err := r.DB.Table("room_schedule").Create(createRoomScheduleDto).Error; err != nil {
		return nil, err
	}

	m := &models.RoomSchedule{
		UserID:    createRoomScheduleDto.UserID,
		RoomID:    createRoomScheduleDto.RoomID,
		StartTime: createRoomScheduleDto.StartTime,
		EndTime:   createRoomScheduleDto.EndTime,
		Status:    models.ScheduleStatus(createRoomScheduleDto.Status),
	}
	return m, nil
}

func (r *RoomScheduleRepository) UpdateRoomSchedule(roomScheduleId int, dto roomSchedule.UpdateRoomSchedule) (*models.RoomSchedule, error) {
	var existingRoomSchedule models.RoomSchedule
	if err := r.DB.Table("room_schedule").Where("id = ?", roomScheduleId).First(&existingRoomSchedule).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"status":     models.ScheduleStatus(dto.Status),
		"room_id":    dto.RoomID,
		"start_time": dto.StartTime,
		"end_time":   dto.EndTime,
		"user_id":    dto.UserID,
	}

	if err := r.DB.Table("room_schedule").Where("id = ?", roomScheduleId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := r.DB.First(&existingRoomSchedule, roomScheduleId).Error; err != nil {
		return nil, err
	}
	return &existingRoomSchedule, nil
}

func (r *RoomScheduleRepository) DeleteRoomSchedule(roomScheduleId int) error {
	result := r.DB.Table("room_schedule").Where("id = ?", roomScheduleId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (r *RoomScheduleRepository) GetAllRoomSchedules() ([]*models.RoomSchedule, error) {
	var roomSchedules []*models.RoomSchedule
	if err := r.DB.Table("room_schedule").Find(&roomSchedules).Error; err != nil {
		return nil, err
	}
	return roomSchedules, nil
}

func NewRoomScheduleRepository(db *gorm.DB) Repository.RoomScheduleRepositoryInterface {
	return &RoomScheduleRepository{
		DB: db,
	}
}
