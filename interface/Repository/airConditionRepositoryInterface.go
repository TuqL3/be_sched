package Repository

import (
	"server/dtos/aircondition"
	"server/models"
)

type AirConditionRepositoryInterface interface {
	CreateAirCondition(createAirConditionDto *aircondition.CreateAirConditionDto) (*models.AirCondition, error)
	UpdateAirCondition(airConditionId int, dto aircondition.UpdateAirConditionDto) (*models.AirCondition, error)
	DeleteAirCondition(airConditionId int) error
	GetAllAirConditions() ([]*models.AirCondition, error)
}
