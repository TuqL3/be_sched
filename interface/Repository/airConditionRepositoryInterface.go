package Repository

import (
	"server/dtos/airCondition"
	"server/models"
)

type AirConditionRepositoryInterface interface {
	CreateAirCondition(createAirConditionDto *airCondition.CreateAirConditionDto) (*models.AirCondition, error)
	UpdateAirCondition(airConditionId int, dto airCondition.UpdateAirConditionDto) (*models.AirCondition, error)
	DeleteAirCondition(airConditionId int) error
	GetAllAirConditions() ([]*models.AirCondition, error)
}
