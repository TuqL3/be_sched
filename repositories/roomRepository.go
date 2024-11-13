package repositories

import (
	"errors"
	"gorm.io/gorm"
	"server/dtos/room"
	"server/interface/Repository"
	"server/models"
	"time"
)

type RoomRepository struct {
	DB *gorm.DB
}

func (r *RoomRepository) GetRoomCount() (int64, error) {
	var count int64
	if err := r.DB.Table("room").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *RoomRepository) GetRoomById(roomId uint) (*models.Room, error) {
	var room models.Room
	if err := r.DB.Table("room").Where("id = ?", roomId).First(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepository) CreateRoom(createRoomDto *room.CreateRoomDto) (*models.Room, error) {
	if err := r.DB.Table("room").Create(createRoomDto).Error; err != nil {
		return nil, err
	}

	m := &models.Room{
		Name:   createRoomDto.Name,
		Status: models.RoomStatus(createRoomDto.Status),
	}

	return m, nil
}

func (r *RoomRepository) UpdateRoom(roomId uint, dto room.UpdateRoomDto) (*models.Room, error) {
	var existingRoom models.Room
	if err := r.DB.Table("room").Where("id = ?", roomId).First(&existingRoom).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"name":   dto.Name,
		"status": dto.Status,
	}
	if err := r.DB.Table("room").Where("id = ?", roomId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := r.DB.First(&existingRoom, roomId).Error; err != nil {
		return nil, err
	}
	return &existingRoom, nil
}

func (r *RoomRepository) DeleteRoom(roomId uint) error {
	result := r.DB.Table("room").Where("id = ?", roomId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (r *RoomRepository) GetAllRooms() ([]*models.Room, error) {
	var room []*models.Room
	if err := r.DB.Find(&room).Error; err != nil {
		return nil, err
	}
	return room, nil
}

func NewRoomRepository(db *gorm.DB) Repository.RoomRepositoryInterface {
	return &RoomRepository{
		DB: db,
	}
}
