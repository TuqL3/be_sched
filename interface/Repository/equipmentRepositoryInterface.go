package Repository

import (
	"server/dtos/equipment"
	"server/models"
)

type EquipmentRepositoryInterface interface {
	CreateEquipment(createEquipmentDto *equipment.CreateEquipmentDto) (*models.Equipment, error)
	UpdateEquipment(equipmentId uint, dto equipment.UpdateEquipmentDto) (*models.Equipment, error)
	DeleteEquipment(equipmentId uint) error
	GetAllEquipments() ([]*models.Equipment, error)
	GetEquipmentById(equipmentId uint) (*models.Equipment, error)
}
