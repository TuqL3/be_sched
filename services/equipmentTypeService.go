package services

import (
	"server/dtos/equipmentType"
	"server/interface/Repository"
	"server/interface/Service"
	"server/models"
)

type EquipmentTypeService struct {
	equipmenttypeRepository Repository.EquipmentTypeRepositoryInterface
}

func (c *EquipmentTypeService) CreateEquipmentType(createEquipmentTypeDto *equipmentType.CreateEquipmentTypeDto) (*models.EquipmentType, error) {
	return c.equipmenttypeRepository.CreateEquipmentType(createEquipmentTypeDto)
}

func (c *EquipmentTypeService) DeleteEquipmentType(equipmenttypeId uint) error {
	return c.equipmenttypeRepository.DeleteEquipmentType(equipmenttypeId)
}

func (c *EquipmentTypeService) GetAllEquipmentTypes() ([]*models.EquipmentType, error) {
	return c.equipmenttypeRepository.GetAllEquipmentTypes()
}

func (c *EquipmentTypeService) GetEquipmentTypeById(equipmenttypeId uint) (*models.EquipmentType, error) {
	return c.equipmenttypeRepository.GetEquipmentTypeById(equipmenttypeId)
}

func (c *EquipmentTypeService) UpdateEquipmentType(equipmenttypeId uint, dto equipmentType.UpdateEquipmentTypeDto) (*models.EquipmentType, error) {
	return c.equipmenttypeRepository.UpdateEquipmentType(equipmenttypeId, dto)
}

func NewEquipmentTypeService(equipmenttypeRepo Repository.EquipmentTypeRepositoryInterface) Service.EquipmentTypeServiceInterface {
	return &EquipmentTypeService{
		equipmenttypeRepository: equipmenttypeRepo,
	}
}
