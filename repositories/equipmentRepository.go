package repositories

import (
	"errors"
	"gorm.io/gorm"
	"server/dtos/equipment"
	"server/interface/Repository"
	"server/models"
	"time"
)

type EquipmentRepository struct {
	DB *gorm.DB
}

func (e *EquipmentRepository) CreateEquipment(createEquipmentDto *equipment.CreateEquipmentDto) (*models.Equipment, error) {
	if err := e.DB.Table("equipment").Create(createEquipmentDto).Error; err != nil {
		return nil, err
	}

	m := &models.Equipment{
		EquipmentName: createEquipmentDto.EquipmentName,
		RoomID:        createEquipmentDto.RoomID,
		Status:        models.EquipmentStatus(createEquipmentDto.Status),
	}
	return m, nil
}

func (e *EquipmentRepository) UpdateEquipment(equipmentId int, dto equipment.UpdateEquipmentDto) (*models.Equipment, error) {
	var existingEquipment models.Equipment
	if err := e.DB.Table("equipment").Where("id = ?", equipmentId).First(&existingEquipment).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"equipment_name": dto.EquipmentName,
		"room_id":        dto.RoomID,
		"status":         dto.Status,
	}

	if err := e.DB.Table("equipment").Where("id = ?", equipmentId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := e.DB.First(&existingEquipment, equipmentId).Error; err != nil {
		return nil, err
	}
	return &existingEquipment, nil
}

func (e *EquipmentRepository) DeleteEquipment(equipmentId int) error {
	result := e.DB.Table("equipment").Where("id = ?", equipmentId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (e *EquipmentRepository) GetAllEquipments() ([]*models.Equipment, error) {
	var equipments []*models.Equipment
	if err := e.DB.Find(&equipments).Error; err != nil {
		return nil, err
	}
	return equipments, nil
}

func NewEquipmentRepository(db *gorm.DB) Repository.EquipmentRepositoryInterface {
	return &EquipmentRepository{
		DB: db,
	}
}
