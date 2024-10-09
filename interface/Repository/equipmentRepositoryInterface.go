package Repository

import (
	"server/dtos/equipment"
	"server/models"
)

type EquipmentRepositoryInterface interface {
	CreateEquipment(createEquipmentDto *equipment.CreateEquipmentDto) (*models.Equipment, error)
	UpdateEquipment(equipmentId int, dto equipment.UpdateEquipmentDto) (*models.Equipment, error)
	DeleteEquipment(equipmentId int) error
	GetAllEquipments() ([]*models.Equipment, error)
}
