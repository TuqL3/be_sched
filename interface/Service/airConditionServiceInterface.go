package Service

import (
	"server/dtos/airCondition"
	"server/models"
)

type AirConditionServiceInterface interface {
	CreateAirCondition(createAirConditionDto *airCondition.CreateAirConditionDto) (*models.AirCondition, error)
	UpdateAirCondition(airConditionId int, dto airCondition.UpdateAirConditionDto) (*models.AirCondition, error)
	DeleteAirCondition(airConditionId int) error
	GetAllAirConditions() ([]*models.AirCondition, error)
}
