package repositories

import (
	"errors"
	"server/dtos/equipment"
	"server/interface/Repository"
	"server/models"
	"server/utils"
	"time"

	"gorm.io/gorm"
)

type EquipmentRepository struct {
	DB *gorm.DB
}

func (a *EquipmentRepository) GetQuantityByStatus() ([]*utils.QuantityStatus, error) {
	var count []*utils.QuantityStatus
	if err := a.DB.Table("equipment").
		Select("status, COUNT(id) as count").
		Group("status").
		Scan(&count).Error; err != nil {
		return nil, err
	}
	return count, nil
}

func (a *EquipmentRepository) GetEquipmentById(equipmentId uint) (*models.Equipment, error) {
	var equipment models.Equipment
	if err := a.DB.Table("equipment").Where("id = ?", equipmentId).Preload("Room").Preload("EquipmentType").First(&equipment).Error; err != nil {
		return nil, err
	}
	return &equipment, nil
}

func (a *EquipmentRepository) CreateEquipment(createEquipmentDto *equipment.CreateEquipmentDto) (*models.Equipment, error) {
	if err := a.DB.Table("equipment").Create(createEquipmentDto).Error; err != nil {
		return nil, err
	}

	m := &models.Equipment{
		Name:            createEquipmentDto.Name,
		RoomID:          createEquipmentDto.RoomID,
		EquipmentTypeID: createEquipmentDto.EquipmentTypeID,
		Status:          utils.EquipmentStatus(createEquipmentDto.Status),
	}
	return m, nil
}

func (a *EquipmentRepository) UpdateEquipment(equipmentId uint, dto equipment.UpdateEquipmentDto) (*models.Equipment, error) {
	var existingEquipment models.Equipment
	if err := a.DB.Table("equipment").Where("id = ?", equipmentId).First(&existingEquipment).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"name":              dto.Name,
		"room_id":           dto.RoomID,
		"status":            dto.Status,
		"equipment_type_id": dto.EquipmentTypeID,
	}

	if err := a.DB.Table("equipment").Where("id = ?", equipmentId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := a.DB.First(&existingEquipment, equipmentId).Error; err != nil {
		return nil, err
	}
	return &existingEquipment, nil
}

func (a *EquipmentRepository) DeleteEquipment(equipmentId uint) error {
	result := a.DB.Table("equipment").Where("id = ?", equipmentId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (a *EquipmentRepository) GetAllEquipments() ([]*models.Equipment, error) {
	var equipments []*models.Equipment
	if err := a.DB.Preload("Room").Preload("EquipmentType").Find(&equipments).Error; err != nil {
		return nil, err
	}
	return equipments, nil
}

func NewEquipmentRepository(db *gorm.DB) Repository.EquipmentRepositoryInterface {
	return &EquipmentRepository{
		DB: db,
	}
}
