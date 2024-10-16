package repositories

import (
	"errors"
	"gorm.io/gorm"
	"server/dtos/aircondition"
	"server/interface/Repository"
	"server/models"
	"time"
)

type AirConditionRepository struct {
	DB *gorm.DB
}

func (a *AirConditionRepository) CreateAirCondition(createAirConditionDto *aircondition.CreateAirConditionDto) (*models.AirCondition, error) {
	if err := a.DB.Table("aircondition").Create(createAirConditionDto).Error; err != nil {
		return nil, err
	}

	m := &models.AirCondition{
		Name:   createAirConditionDto.Name,
		RoomID: createAirConditionDto.RoomID,
		Status: models.EquipmentStatus(createAirConditionDto.Status),
	}
	return m, nil
}

func (a *AirConditionRepository) UpdateAirCondition(airConditionId int, dto aircondition.UpdateAirConditionDto) (*models.AirCondition, error) {
	var existingAirCondition models.AirCondition
	if err := a.DB.Table("aircondition").Where("id = ?", airConditionId).First(&existingAirCondition).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"name":    dto.Name,
		"room_id": dto.RoomID,
		"status":  dto.Status,
	}

	if err := a.DB.Table("aircondition").Where("id = ?", airConditionId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := a.DB.First(&existingAirCondition, airConditionId).Error; err != nil {
		return nil, err
	}
	return &existingAirCondition, nil
}

func (a *AirConditionRepository) DeleteAirCondition(airConditionId int) error {
	result := a.DB.Table("aircondition").Where("id = ?", airConditionId).Update("deleted_at", time.Now())
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
	if err := a.DB.Find(&airConditions).Error; err != nil {
		return nil, err
	}
	return airConditions, nil
}

func NewAirConditionRepository(db *gorm.DB) Repository.AirConditionRepositoryInterface {
	return &AirConditionRepository{
		DB: db,
	}
}
