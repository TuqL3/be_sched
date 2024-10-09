package services

import (
	"server/dtos/equipment"
	"server/interface/Repository"
	"server/interface/Service"
	"server/models"
)

type EquipmentService struct {
	equipmentService Repository.EquipmentRepositoryInterface
}

func (e *EquipmentService) CreateEquipment(createEquipmentDto *equipment.CreateEquipmentDto) (*models.Equipment, error) {
	return e.equipmentService.CreateEquipment(createEquipmentDto)
}

func (e *EquipmentService) UpdateEquipment(equipmentId int, dto equipment.UpdateEquipmentDto) (*models.Equipment, error) {
	return e.equipmentService.UpdateEquipment(equipmentId, dto)
}

func (e *EquipmentService) DeleteEquipment(equipmentId int) error {
	return e.equipmentService.DeleteEquipment(equipmentId)
}

func (e *EquipmentService) GetAllEquipments() ([]*models.Equipment, error) {
	return e.equipmentService.GetAllEquipments()
}

func NewEquipmentService(equipment Repository.EquipmentRepositoryInterface) Service.EquipmentServiceInterface {
	return &EquipmentService{
		equipmentService: equipment,
	}
}
