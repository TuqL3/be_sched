package repositories

import (
	"errors"
	"server/dtos/schedule"
	"server/interface/Repository"
	"server/models"
	"server/utils"
	"time"

	"gorm.io/gorm"
)

type ScheduleRepository struct {
	DB *gorm.DB
}

func (r *ScheduleRepository) CreateSchedule(createScheduleDto *schedule.CreateRoomScheduleDto) (*models.Schedule, error) {
	if err := r.DB.Table("schedule").Create(createScheduleDto).Error; err != nil {
		return nil, err
	}

	m := &models.Schedule{
		UserID:    createScheduleDto.UserID,
		RoomID:    createScheduleDto.RoomID,
		StartTime: createScheduleDto.StartTime,
		EndTime:   createScheduleDto.EndTime,
		Status:    utils.ScheduleStatus(createScheduleDto.Status),
	}
	return m, nil
}

func (r *ScheduleRepository) UpdateSchedule(roomScheduleId uint, dto schedule.UpdateRoomSchedule) (*models.Schedule, error) {
	var existingSchedule models.Schedule
	if err := r.DB.Table("schedule").Where("id = ?", roomScheduleId).First(&existingSchedule).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"status":     utils.ScheduleStatus(dto.Status),
		"room_id":    dto.RoomID,
		"start_time": dto.StartTime,
		"end_time":   dto.EndTime,
		"user_id":    dto.UserID,
	}

	if err := r.DB.Table("schedule").Where("id = ?", roomScheduleId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := r.DB.First(&existingSchedule, roomScheduleId).Error; err != nil {
		return nil, err
	}
	return &existingSchedule, nil
}

func (r *ScheduleRepository) DeleteSchedule(roomScheduleId uint) error {
	result := r.DB.Table("schedule").Where("id = ?", roomScheduleId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (r *ScheduleRepository) GetAllSchedules() ([]*models.Schedule, error) {
	var roomSchedules []*models.Schedule
	if err := r.DB.Table("schedule").Find(&roomSchedules).Error; err != nil {
		return nil, err
	}
	return roomSchedules, nil
}

func NewScheduleRepository(db *gorm.DB) Repository.RoomScheduleRepositoryInterface {
	return &ScheduleRepository{
		DB: db,
	}
}
