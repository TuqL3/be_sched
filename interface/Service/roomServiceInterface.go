package Service

import (
	"server/dtos/room"
	"server/models"
)

type RoomServiceInterface interface {
	CreateRoom(createRoomDto *room.CreateRoomDto) (*models.Room, error)
	UpdateRoom(roomId uint, dto room.UpdateRoomDto) (*models.Room, error)
	DeleteRoom(roomId uint) error
	GetAllRooms() ([]*models.Room, error)
	GetRoomById(roomId uint) (*models.Room, error)
	GetRoomCount() (int64, error)
}
