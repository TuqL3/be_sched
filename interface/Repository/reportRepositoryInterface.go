package Repository

import (
	"server/dtos/report"
	"server/models"
)

type ReportRepositoryInterface interface {
	CreateReport(createReportDto *report.CreateReportDto) (*models.Report, error)
	UpdateReport(reportId int, dto report.UpdateReportDto) (*models.Report, error)
	DeleteReport(reportId int) error
	GetAllReports() ([]*models.Report, error)
}
