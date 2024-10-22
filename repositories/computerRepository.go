package repositories

import (
	"errors"
	"gorm.io/gorm"
	"server/dtos/computer"
	"server/interface/Repository"
	"server/models"
	"server/utils"
	"time"
)

type ComputerRepository struct {
	DB *gorm.DB
}

func (e *ComputerRepository) CreateCompute(createComputeDto *computer.CreateComputerDto) (*models.Computer, error) {
	if err := e.DB.Table("computer").Create(createComputeDto).Error; err != nil {
		return nil, err
	}

	m := &models.Computer{
		Name:   createComputeDto.Name,
		RoomID: createComputeDto.RoomID,
		Status: utils.EquipmentStatus(createComputeDto.Status),
	}
	return m, nil
}

func (e *ComputerRepository) UpdateCompute(computerId int, dto computer.UpdateComputerDto) (*models.Computer, error) {
	var existingComputer models.Computer
	if err := e.DB.Table("computer").Where("id = ?", computerId).First(&existingComputer).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"name":    dto.Name,
		"room_id": dto.RoomID,
		"status":  dto.Status,
	}

	if err := e.DB.Table("computer").Where("id = ?", computerId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := e.DB.First(&existingComputer, computerId).Error; err != nil {
		return nil, err
	}
	return &existingComputer, nil
}

func (e *ComputerRepository) DeleteCompute(computerId int) error {
	result := e.DB.Table("computer").Where("id = ?", computerId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (e *ComputerRepository) GetAllComputes() ([]*models.Computer, error) {
	var computers []*models.Computer
	if err := e.DB.Find(&computers).Error; err != nil {
		return nil, err
	}
	return computers, nil
}

func NewComputerRepository(db *gorm.DB) Repository.ComputerRepositoryInterface {
	return &ComputerRepository{
		DB: db,
	}
}
