package Service

import (
	"server/dtos/tandch"
	"server/models"
)

type TAndChServiceInterface interface {
	CreateTAndCh(createTAndChDto *tandch.CreateTandChDto) (*models.TandCh, error)
	UpdateTAndCh(TAndChId int, dto tandch.UpdateTandChDto) (*models.TandCh, error)
	DeleteTAndCh(TAndChId int) error
	GetAllTAndChs() ([]*models.TandCh, error)
}
