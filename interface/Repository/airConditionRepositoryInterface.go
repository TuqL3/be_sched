package Repository

import (
	"server/dtos/airCondition"
	"server/models"
)

type AirConditionRepositoryInterface interface {
	CreateAirCondition(createAirConditionDto *airCondition.CreateAirConditionDto) (*models.AirCondition, error)
	UpdateAirCondition(airConditionId uint, dto airCondition.UpdateAirConditionDto) (*models.AirCondition, error)
	DeleteAirCondition(airConditionId uint) error
	GetAllAirConditions() ([]*models.AirCondition, error)
	GetAirConditionById(airConditionId uint) (*models.AirCondition, error)
}
