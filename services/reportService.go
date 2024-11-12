package services

import (
	"server/dtos/report"
	"server/interface/Repository"
	"server/interface/Service"
	"server/models"
	"server/utils"
)

type ReportService struct {
	reportRepository Repository.ReportRepositoryInterface
}

func (r *ReportService) GetCountReport() (int64, error) {
	return r.reportRepository.GetCountReport()
}

func (r *ReportService) GetQuantityReportOfRoom() ([]*utils.ReportRoomCount, error) {
	return r.reportRepository.GetQuantityReportOfRoom()
}

func (r *ReportService) GetReportById(reportId uint) (*models.Report, error) {
	return r.reportRepository.GetReportById(reportId)
}

func (r *ReportService) CreateReport(createReportDto *report.CreateReportDto) (*models.Report, error) {
	return r.reportRepository.CreateReport(createReportDto)
}

func (r *ReportService) UpdateReport(reportId uint, dto report.UpdateReportDto) (*models.Report, error) {
	return r.reportRepository.UpdateReport(reportId, dto)
}

func (r *ReportService) DeleteReport(reportId uint) error {
	return r.reportRepository.DeleteReport(reportId)
}

func (r *ReportService) GetAllReports() ([]*models.Report, error) {
	return r.reportRepository.GetAllReports()
}

func NewReportService(reportRepository Repository.ReportRepositoryInterface) Service.ReportServiceInterface {
	return &ReportService{
		reportRepository: reportRepository,
	}
}
