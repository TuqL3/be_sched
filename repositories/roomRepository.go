package repositories

import (
	"errors"
	"server/dtos/room"
	"server/interface/Repository"
	"server/models"
	"time"

	"gorm.io/gorm"
)

type RoomRepository struct {
	DB *gorm.DB
}

func (r *RoomRepository) CreateRoom(createRoomDto *room.CreateRoomDto) (*models.Room, error) {
	if err := r.DB.Table("rooms").Create(createRoomDto).Error; err != nil {
		return nil, err
	}

	m := &models.Room{
		RoomName: createRoomDto.RoomName,
		Capacity: createRoomDto.Capacity,
		Status:   models.RoomStatus(createRoomDto.Status),
	}

	return m, nil
}

func (r *RoomRepository) UpdateRoom(roomId uint, dto room.UpdateRoomDto) (*models.Room, error) {
	var existingRoom models.Room
	if err := r.DB.Table("rooms").Where("id = ?", roomId).First(&existingRoom).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"room_name": dto.RoomName,
		"capacity":  dto.Capacity,
		"status":    dto.Status,
	}
	if err := r.DB.Table("rooms").Where("id = ?", roomId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := r.DB.First(&existingRoom, roomId).Error; err != nil {
		return nil, err
	}
	return &existingRoom, nil
}

func (r *RoomRepository) DeleteRoom(roomId uint) error {
	result := r.DB.Table("rooms").Where("id = ?", roomId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (r *RoomRepository) GetAllRooms() ([]*models.Room, error) {
	var rooms []*models.Room
	if err := r.DB.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

func NewRoomRepository(db *gorm.DB) Repository.RoomRepositoryInterface {
	return &RoomRepository{
		DB: db,
	}
}
