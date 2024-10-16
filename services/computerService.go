package services

import (
	"server/dtos/computer"
	"server/interface/Repository"
	"server/interface/Service"
	"server/models"
)

type ComputerService struct {
	computerService Repository.ComputerRepositoryInterface
}

func (c *ComputerService) CreateCompute(createComputeDto *computer.CreateComputerDto) (*models.Computer, error) {
	return c.computerService.CreateCompute(createComputeDto)
}

func (c *ComputerService) UpdateCompute(computerId int, dto computer.UpdateComputerDto) (*models.Computer, error) {
	return c.computerService.UpdateCompute(computerId, dto)
}

func (c *ComputerService) DeleteCompute(computerId int) error {
	return c.computerService.DeleteCompute(computerId)
}

func (c *ComputerService) GetAllComputes() ([]*models.Computer, error) {
	return c.computerService.GetAllComputes()
}

func NewComputerService(computerCondition Repository.ComputerRepositoryInterface) Service.ComputerServiceInterface {
	return &ComputerService{
		computerService: computerCondition,
	}
}
