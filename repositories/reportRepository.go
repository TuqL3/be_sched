package repositories

import (
	"errors"
	"server/dtos/report"
	"server/interface/Repository"
	"server/models"
	"time"

	"gorm.io/gorm"
)

type ReportRepository struct {
	DB *gorm.DB
}

func (r *ReportRepository) GetReportById(reportId uint) (*models.Report, error) {
	var report models.Report
	if err := r.DB.Table("report").Where("id = ?", reportId).Preload("Room").First(&report).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *ReportRepository) CreateReport(createReportDto *report.CreateReportDto) (*models.Report, error) {
	if err := r.DB.Table("reports").Create(createReportDto).Error; err != nil {
		return nil, err
	}

	m := &models.Report{
		UserID:        createReportDto.UserID,
		RoomID:        createReportDto.RoomID,
		EquipmentID:   createReportDto.EquipmentID,
		EquipmentType: createReportDto.EquipmentType,
		Description:   createReportDto.Description,
		Status:        models.ReportStatus(createReportDto.Status),
	}
	return m, nil
}

func (r *ReportRepository) UpdateReport(reportId uint, dto report.UpdateReportDto) (*models.Report, error) {
	var existingReport models.Report
	if err := r.DB.Table("reports").Where("id = ?", reportId).First(&existingReport).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"description":    dto.Description,
		"status":         dto.Status,
		"room_id":        dto.RoomID,
		"user_id":        dto.UserID,
		"equipment_id":   dto.EquipmentID,
		"equipment_type": dto.EquipmentType,
	}
	if err := r.DB.Table("reports").Where("id = ?", reportId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := r.DB.First(&existingReport, reportId).Error; err != nil {
		return nil, err
	}
	return &existingReport, nil
}

func (r *ReportRepository) DeleteReport(reportId uint) error {
	result := r.DB.Table("reports").Where("id = ?", reportId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("report not found")
	}
	return nil
}

func (r *ReportRepository) GetAllReports() ([]*models.Report, error) {
	var reports []*models.Report
	if err := r.DB.Find(&reports).Error; err != nil {
		return nil, err
	}
	return reports, nil
}

func NewReportRepository(db *gorm.DB) Repository.ReportRepositoryInterface {
	return &ReportRepository{
		DB: db,
	}
}
