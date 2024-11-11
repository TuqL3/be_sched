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

func (r *ScheduleRepository) GetcountScheduleUser() ([]*utils.ScheduleUserCount, error) {
	var counts []*utils.ScheduleUserCount
	if err := r.DB.Table("schedule").
		Select("user.id as user_id, user.name as user_name, COUNT(schedule.id) as schedule_count").
		Group("user.id, user.name").
		Scan(&counts).Error; err != nil {
		return nil, err
	}
	return counts, nil
}

func (r *ScheduleRepository) GetCountScheduleRoom() ([]*utils.ScheduleRoomCount, error) {
	var counts []*utils.ScheduleRoomCount
	if err := r.DB.Table("schedule").
		Select("room.id as room_id, room.name as room_name, COUNT(schedule.id) as schedule_count").
		Joins("JOIN room ON schedule.room_id = room.id").
		Group("room.id, room.name").
		Scan(&counts).Error; err != nil {
		return nil, err
	}
	return counts, nil

}

func (r *ScheduleRepository) CreateSchedule(createScheduleDto *schedule.CreateRoomScheduleDto) (*models.Schedule, error) {
	newSchedule := &models.Schedule{
		UserID:      createScheduleDto.UserID,
		RoomID:      createScheduleDto.RoomID,
		StartTime:   createScheduleDto.StartTime,
		EndTime:     createScheduleDto.EndTime,
		Status:      utils.ScheduleStatus(createScheduleDto.Status),
		Description: createScheduleDto.Description,
		Title:       createScheduleDto.Title,
	}

	if err := r.DB.Create(newSchedule).Error; err != nil {
		return nil, err
	}

	return newSchedule, nil
}

func (r *ScheduleRepository) UpdateSchedule(roomScheduleId uint, dto schedule.UpdateRoomSchedule) (*models.Schedule, error) {
	var existingSchedule models.Schedule
	if err := r.DB.Table("schedule").Where("id = ?", roomScheduleId).First(&existingSchedule).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"status":      utils.ScheduleStatus(dto.Status),
		"room_id":     dto.RoomID,
		"start_time":  dto.StartTime,
		"end_time":    dto.EndTime,
		"user_id":     dto.UserID,
		"description": dto.Description,
		"title":       dto.Title,
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
	if err := r.DB.Table("schedule").Preload("User").Preload("Room").Find(&roomSchedules).Error; err != nil {
		return nil, err
	}
	return roomSchedules, nil
}

func NewScheduleRepository(db *gorm.DB) Repository.RoomScheduleRepositoryInterface {
	return &ScheduleRepository{
		DB: db,
	}
}
