package Service

import (
	"server/dtos/equipmentType"
	"server/models"
)

type EquipmentTypeServiceInterface interface {
	CreateEquipmentType(createEquipmentTypeDto *equipmentType.CreateEquipmentTypeDto) (*models.EquipmentType, error)
	UpdateEquipmentType(equipmentTypeId uint, dto equipmentType.UpdateEquipmentTypeDto) (*models.EquipmentType, error)
	DeleteEquipmentType(equipmentTypeId uint) error
	GetAllEquipmentTypes() ([]*models.EquipmentType, error)
	GetEquipmentTypeById(equipmentTypeId uint) (*models.EquipmentType, error)
}
