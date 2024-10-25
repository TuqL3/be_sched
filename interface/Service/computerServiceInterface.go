package Service

import (
	"server/dtos/computer"
	"server/models"
)

type ComputerServiceInterface interface {
	CreateCompute(createComputeDto *computer.CreateComputerDto) (*models.Computer, error)
	UpdateCompute(computerId uint, dto computer.UpdateComputerDto) (*models.Computer, error)
	DeleteCompute(computerId uint) error
	GetComputerById(computerId uint) (*models.Computer, error)
	GetAllComputes() ([]*models.Computer, error)
}
