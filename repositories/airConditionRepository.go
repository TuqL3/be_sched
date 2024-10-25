package repositories

import (
	"errors"
	"server/dtos/airCondition"
	"server/interface/Repository"
	"server/models"
	"server/utils"
	"time"

	"gorm.io/gorm"
)

type AirConditionRepository struct {
	DB *gorm.DB
}

func (a *AirConditionRepository) GetAirConditionById(airConditionId uint) (*models.AirCondition, error) {
	var airCondition models.AirCondition
	if err := a.DB.Table("airCondition").Where("id = ?", airConditionId).Preload("Room").First(&airCondition).Error; err != nil {
		return nil, err
	}
	return &airCondition, nil
}

func (a *AirConditionRepository) CreateAirCondition(createAirConditionDto *airCondition.CreateAirConditionDto) (*models.AirCondition, error) {
	if err := a.DB.Table("airCondition").Create(createAirConditionDto).Error; err != nil {
		return nil, err
	}

	m := &models.AirCondition{
		Name:   createAirConditionDto.Name,
		RoomID: createAirConditionDto.RoomID,
		Status: utils.EquipmentStatus(createAirConditionDto.Status),
	}
	return m, nil
}

func (a *AirConditionRepository) UpdateAirCondition(airConditionId uint, dto airCondition.UpdateAirConditionDto) (*models.AirCondition, error) {
	var existingAirCondition models.AirCondition
	if err := a.DB.Table("airCondition").Where("id = ?", airConditionId).First(&existingAirCondition).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"name":    dto.Name,
		"room_id": dto.RoomID,
		"status":  dto.Status,
	}

	if err := a.DB.Table("airCondition").Where("id = ?", airConditionId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := a.DB.First(&existingAirCondition, airConditionId).Error; err != nil {
		return nil, err
	}
	return &existingAirCondition, nil
}

func (a *AirConditionRepository) DeleteAirCondition(airConditionId uint) error {
	result := a.DB.Table("airCondition").Where("id = ?", airConditionId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (a *AirConditionRepository) GetAllAirConditions() ([]*models.AirCondition, error) {
	var airConditions []*models.AirCondition
	if err := a.DB.Preload("Room").Find(&airConditions).Error; err != nil {
		return nil, err
	}
	return airConditions, nil
}

func NewAirConditionRepository(db *gorm.DB) Repository.AirConditionRepositoryInterface {
	return &AirConditionRepository{
		DB: db,
	}
}
