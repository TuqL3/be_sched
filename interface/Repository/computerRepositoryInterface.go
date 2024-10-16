package Repository

import (
	"server/dtos/computer"
	"server/models"
)

type ComputerRepositoryInterface interface {
	CreateCompute(createComputeDto *computer.CreateComputerDto) (*models.Computer, error)
	UpdateCompute(computerId int, dto computer.UpdateComputerDto) (*models.Computer, error)
	DeleteCompute(computerId int) error
	GetAllComputes() ([]*models.Computer, error)
}
