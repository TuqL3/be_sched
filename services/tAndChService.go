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

func (T *TAndChService) GetTAndChById(TAndChId uint) (*models.TandCh, error) {
	return T.tAndChService.GetTAndChById(TAndChId)
}

func (T *TAndChService) CreateTAndCh(createTAndChDto *tandch.CreateTandChDto) (*models.TandCh, error) {
	return T.tAndChService.CreateTAndCh(createTAndChDto)
}

func (T *TAndChService) UpdateTAndCh(TAndChId uint, dto tandch.UpdateTandChDto) (*models.TandCh, error) {
	return T.tAndChService.UpdateTAndCh(TAndChId, dto)
}

func (T *TAndChService) DeleteTAndCh(TAndChId uint) error {
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
