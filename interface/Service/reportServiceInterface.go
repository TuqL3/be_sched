package Service

import (
	"server/dtos/report"
	"server/models"
)

type ReportServiceInterface interface {
	CreateReport(createReportDto *report.CreateReportDto) (*models.Report, error)
	UpdateReport(reportId uint, dto report.UpdateReportDto) (*models.Report, error)
	DeleteReport(reportId uint) error
	GetAllReports() ([]*models.Report, error)
	GetReportById(reportId uint) (*models.Report, error)
}
