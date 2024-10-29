package repositories

import (
	"errors"
	"server/dtos/equipmentType"
	"server/interface/Repository"
	"server/models"
	"time"

	"gorm.io/gorm"
)

type EquipmentTypeRepository struct {
	DB *gorm.DB
}

func (c *EquipmentTypeRepository) CreateEquipmentType(createEquipmentTypeDto *equipmentType.CreateEquipmentTypeDto) (*models.EquipmentType, error) {
	if err := c.DB.Table("equipmentType").Create(createEquipmentTypeDto).Error; err != nil {
		return nil, err
	}
	m := &models.EquipmentType{
		Name: createEquipmentTypeDto.Name,
	}

	return m, nil
}

func (c *EquipmentTypeRepository) DeleteEquipmentType(equipmenttypeId uint) error {
	result := c.DB.Table("equipmentType").Where("id = ?", equipmenttypeId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (c *EquipmentTypeRepository) GetAllEquipmentTypes() ([]*models.EquipmentType, error) {
	var equipmenttype []*models.EquipmentType
	if err := c.DB.Find(&equipmenttype).Error; err != nil {
		return nil, err
	}
	return equipmenttype, nil
}

func (c *EquipmentTypeRepository) GetEquipmentTypeById(equipmenttypeId uint) (*models.EquipmentType, error) {
	var equipmenttype models.EquipmentType
	if err := c.DB.Table("equipmentType").Where("id = ?", equipmenttypeId).First(&equipmenttype).Error; err != nil {
		return nil, err
	}
	return &equipmenttype, nil
}

func (c *EquipmentTypeRepository) UpdateEquipmentType(equipmenttypeId uint, dto equipmentType.UpdateEquipmentTypeDto) (*models.EquipmentType, error) {
	var existingEquipmentType models.EquipmentType
	if err := c.DB.Table("equipmentType").Where("id = ?", equipmenttypeId).First(&existingEquipmentType).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"name": dto.Name,
	}
	if err := c.DB.Table("equipmentType").Where("id = ?", equipmenttypeId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := c.DB.First(&existingEquipmentType, equipmenttypeId).Error; err != nil {
		return nil, err
	}
	return &existingEquipmentType, nil
}

func NewEquipmentTypeRepository(db *gorm.DB) Repository.EquipmentTypeRepositoryInterface {
	return &EquipmentTypeRepository{
		DB: db,
	}
}
