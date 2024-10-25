package Repository

import (
	"server/dtos/computer"
	"server/models"
)

type ComputerRepositoryInterface interface {
	CreateCompute(createComputeDto *computer.CreateComputerDto) (*models.Computer, error)
	UpdateCompute(computerId uint, dto computer.UpdateComputerDto) (*models.Computer, error)
	DeleteCompute(computerId uint) error
	GetAllComputes() ([]*models.Computer, error)
	GetComputerById(computerId uint) (*models.Computer, error)
}
