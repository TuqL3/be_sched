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

func (e *EquipmentService) GetEquipmentById(equipmentId uint) (*models.Equipment, error) {
	return e.equipmentService.GetEquipmentById(equipmentId)
}

func (e *EquipmentService) CreateEquipment(createEquipmentDto *equipment.CreateEquipmentDto) (*models.Equipment, error) {
	return e.equipmentService.CreateEquipment(createEquipmentDto)
}

func (e *EquipmentService) UpdateEquipment(equipmentId uint, dto equipment.UpdateEquipmentDto) (*models.Equipment, error) {
	return e.equipmentService.UpdateEquipment(equipmentId, dto)
}

func (e *EquipmentService) DeleteEquipment(equipmentId uint) error {
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
