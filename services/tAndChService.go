package services

import (
	"server/dtos/tandch"
	"server/interface/Repository"
	"server/interface/Service"
	"server/models"
)

type TAndChService struct {
	tAndChService Repository.TAndChRepositoryInterface
}

func (T *TAndChService) CreateTAndCh(createTAndChDto *tandch.CreateTandChDto) (*models.TandCh, error) {
	return T.tAndChService.CreateTAndCh(createTAndChDto)
}

func (T *TAndChService) UpdateTAndCh(TAndChId int, dto tandch.UpdateTandChDto) (*models.TandCh, error) {
	return T.tAndChService.UpdateTAndCh(TAndChId, dto)
}

func (T *TAndChService) DeleteTAndCh(TAndChId int) error {
	return T.tAndChService.DeleteTAndCh(TAndChId)
}

func (T *TAndChService) GetAllTAndChs() ([]*models.TandCh, error) {
	return T.tAndChService.GetAllTAndChs()
}

func NewTAndChService(tAndCh Repository.TAndChRepositoryInterface) Service.TAndChServiceInterface {
	return &TAndChService{
		tAndChService: tAndCh,
	}
}
