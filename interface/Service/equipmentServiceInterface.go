package Service

import (
	"server/dtos/equipment"
	"server/models"
	"server/utils"
)

type EquipmentServiceInterface interface {
	CreateEquipment(createEquipmentDto *equipment.CreateEquipmentDto) (*models.Equipment, error)
	UpdateEquipment(equipmentId uint, dto equipment.UpdateEquipmentDto) (*models.Equipment, error)
	DeleteEquipment(equipmentId uint) error
	GetEquipmentById(equipmentId uint) (*models.Equipment, error)
	GetAllEquipments() ([]*models.Equipment, error)
	GetQuantityByStatus() ([]*utils.QuantityStatus, error)
	GetCountEquipment() (int64, error)
}
