package Repository

import (
	"server/dtos/room"
	"server/models"
)

type RoomRepositoryInterface interface {
	CreateRoom(createRoomDto *room.CreateRoomDto) (*models.Room, error)
	UpdateRoom(roomId uint, dto room.UpdateRoomDto) (*models.Room, error)
	DeleteRoom(roomId uint) error
	GetAllRooms() ([]*models.Room, error)
	GetRoomById(roomId uint) (*models.Room, error)
}
