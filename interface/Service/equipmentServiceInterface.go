package Service

import (
	"server/dtos/equipment"
	"server/models"
)

type EquipmentServiceInterface interface {
	CreateEquipment(createEquipmentDto *equipment.CreateEquipmentDto) (*models.Equipment, error)
	UpdateEquipment(equipmentId int, dto equipment.UpdateEquipmentDto) (*models.Equipment, error)
	DeleteEquipment(equipmentId int) error
	GetAllEquipments() ([]*models.Equipment, error)
}
