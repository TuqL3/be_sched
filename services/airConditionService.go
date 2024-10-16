package services

import (
	"server/dtos/airCondition"
	"server/interface/Repository"
	"server/interface/Service"
	"server/models"
)

type AirConditionService struct {
	airConditionService Repository.AirConditionRepositoryInterface
}

func (e *AirConditionService) CreateAirCondition(createAirConditionDto *airCondition.CreateAirConditionDto) (*models.AirCondition, error) {
	return e.airConditionService.CreateAirCondition(createAirConditionDto)
}

func (e *AirConditionService) UpdateAirCondition(airConditionId int, dto airCondition.UpdateAirConditionDto) (*models.AirCondition, error) {
	return e.airConditionService.UpdateAirCondition(airConditionId, dto)
}

func (e *AirConditionService) DeleteAirCondition(airConditionId int) error {
	return e.airConditionService.DeleteAirCondition(airConditionId)
}

func (e *AirConditionService) GetAllAirConditions() ([]*models.AirCondition, error) {
	return e.airConditionService.GetAllAirConditions()
}

func NewAirConditionService(airCondition Repository.AirConditionRepositoryInterface) Service.AirConditionServiceInterface {
	return &AirConditionService{
		airConditionService: airCondition,
	}
}
