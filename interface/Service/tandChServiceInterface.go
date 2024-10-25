package Service

import (
	"server/dtos/tandch"
	"server/models"
)

type TAndChServiceInterface interface {
	CreateTAndCh(createTAndChDto *tandch.CreateTandChDto) (*models.TandCh, error)
	UpdateTAndCh(TAndChId uint, dto tandch.UpdateTandChDto) (*models.TandCh, error)
	DeleteTAndCh(TAndChId uint) error
	GetAllTAndChs() ([]*models.TandCh, error)
	GetTAndChById(TAndChId uint) (*models.TandCh, error)
}
