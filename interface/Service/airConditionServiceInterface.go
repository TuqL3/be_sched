package Service

import (
	"server/dtos/airCondition"
	"server/models"
)

type AirConditionServiceInterface interface {
	CreateAirCondition(createAirConditionDto *airCondition.CreateAirConditionDto) (*models.AirCondition, error)
	UpdateAirCondition(airConditionId uint, dto airCondition.UpdateAirConditionDto) (*models.AirCondition, error)
	DeleteAirCondition(airConditionId uint) error
	GetAirConditionById(airConditionId uint) (*models.AirCondition, error)
	GetAllAirConditions() ([]*models.AirCondition, error)
}
