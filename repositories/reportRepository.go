package repositories

import (
	"errors"
	"server/dtos/report"
	"server/interface/Repository"
	"server/models"
	"server/utils"
	"time"

	"gorm.io/gorm"
)

type ReportRepository struct {
	DB *gorm.DB
}

func (r *ReportRepository) GetQuantityReportOfRoom() ([]*utils.ReportRoomCount, error) {
	var counts []*utils.ReportRoomCount
	if err := r.DB.Table("report").
		Select("room.id as room_id, room.name as room_name, COUNT(report.id) as report_count").
		Joins("JOIN room ON report.room_id = room.id").
		Group("room.id, room.name").
		Scan(&counts).Error; err != nil {
		return nil, err
	}
	return counts, nil
}

func (r *ReportRepository) GetReportById(reportId uint) (*models.Report, error) {
	var report models.Report
	if err := r.DB.Table("report").Where("id = ?", reportId).Preload("Room").First(&report).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *ReportRepository) CreateReport(createReportDto *report.CreateReportDto) (*models.Report, error) {
	if err := r.DB.Table("report").Create(createReportDto).Error; err != nil {
		return nil, err
	}

	m := &models.Report{
		UserID:      createReportDto.UserID,
		RoomID:      createReportDto.RoomID,
		EquipmentID: createReportDto.EquipmentID,
		Content:     createReportDto.Content,
		Status:      utils.ReportStatus(createReportDto.Status),
	}
	return m, nil
}

func (r *ReportRepository) UpdateReport(reportId uint, dto report.UpdateReportDto) (*models.Report, error) {
	var existingReport models.Report
	if err := r.DB.Table("report").Where("id = ?", reportId).First(&existingReport).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"content":      dto.Content,
		"status":       dto.Status,
		"room_id":      dto.RoomID,
		"user_id":      dto.UserID,
		"equipment_id": dto.EquipmentID,
	}
	if err := r.DB.Table("report").Where("id = ?", reportId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := r.DB.First(&existingReport, reportId).Error; err != nil {
		return nil, err
	}
	return &existingReport, nil
}

func (r *ReportRepository) DeleteReport(reportId uint) error {
	result := r.DB.Table("report").Where("id = ?", reportId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("report not found")
	}
	return nil
}

func (r *ReportRepository) GetAllReports() ([]*models.Report, error) {
	var report []*models.Report
	if err := r.DB.Preload("Room").Find(&report).Error; err != nil {
		return nil, err
	}
	return report, nil
}

func NewReportRepository(db *gorm.DB) Repository.ReportRepositoryInterface {
	return &ReportRepository{
		DB: db,
	}
}
