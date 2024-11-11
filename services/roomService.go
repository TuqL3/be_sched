package services

import (
	"server/dtos/room"
	"server/interface/Repository"
	"server/interface/Service"
	"server/models"
)

type RoomService struct {
	roomRepository Repository.RoomRepositoryInterface
}

func (r *RoomService) GetRoomCount() (int64, error) {
	return r.roomRepository.GetRoomCount()
}

func (r *RoomService) GetRoomById(roomId uint) (*models.Room, error) {
	return r.roomRepository.GetRoomById(roomId)
}

func (r *RoomService) CreateRoom(createRoomDto *room.CreateRoomDto) (*models.Room, error) {
	return r.roomRepository.CreateRoom(createRoomDto)
}

func (r *RoomService) UpdateRoom(roomId uint, dto room.UpdateRoomDto) (*models.Room, error) {
	return r.roomRepository.UpdateRoom(roomId, dto)
}

func (r *RoomService) DeleteRoom(roomId uint) error {
	return r.roomRepository.DeleteRoom(roomId)
}

func (r *RoomService) GetAllRooms() ([]*models.Room, error) {
	return r.roomRepository.GetAllRooms()
}

func NewRoomService(roomRepo Repository.RoomRepositoryInterface) Service.RoomServiceInterface {
	return &RoomService{
		roomRepository: roomRepo,
	}
}
