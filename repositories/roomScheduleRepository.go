package repositories

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"server/dtos/schedule"
	"server/interface/Repository"
	"server/models"
	"server/utils"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type ScheduleRepository struct {
	DB *gorm.DB
}

func (r *ScheduleRepository) ImportScheduleFromExcel(file string) error {
	f, err := excelize.OpenFile(file)
	if err != nil {
		return errors.New("failed to open Excel file")
	}

	// Read the first sheet
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return errors.New("failed to read Excel rows")
	}

	// Log toàn bộ dữ liệu trong file Excel
	fmt.Printf("Logging all data from sheet '%s':\n", sheetName)
	for i, row := range rows {
		fmt.Printf("Row %d: %v\n", i+1, row) // Log từng hàng
	}

	// Xử lý dữ liệu từng hàng
	for i, row := range rows {
		if i == 0 {
			continue // Skip header row
		}

		if len(row) < 7 {
			fmt.Printf("Skipping invalid row %d: %v\n", i+1, row)
			continue // Skip invalid rows
		}

		// Log the row data
		fmt.Printf("Processing row %d: %v\n", i+1, row)

		// Parse data
		roomID, _ := strconv.Atoi(row[0])
		userID, _ := strconv.Atoi(row[1])

		// Parse start time
		startTime, err := time.Parse("01/02/06 15:04", row[3])
		if err != nil {
			fmt.Printf("Row %d: invalid start time format: %s\n", i+1, row[3])
			return fmt.Errorf("row %d: invalid start time format: %s", i+1, row[3])
		}
		startTimeUTC := startTime.UTC() // Convert to UTC

		// Parse end time
		endTime, err := time.Parse("01/02/06 15:04", row[4])
		if err != nil {
			fmt.Printf("Row %d: invalid end time format: %s\n", i+1, row[4])
			return fmt.Errorf("row %d: invalid end time format: %s", i+1, row[4])
		}
		endTimeUTC := endTime.UTC() // Convert to UTC

		// Create schedule entity
		schedule := models.Schedule{
			RoomID:      uint(roomID),
			UserID:      uint(userID),
			StartTime:   startTimeUTC,
			EndTime:     endTimeUTC,
			Status:      models.ScheduleStatus(row[5]),
			Description: row[6],
			Title:       row[2],
		}

		// Log schedule object before saving
		fmt.Printf("Row %d: schedule created: %+v\n", i+1, schedule)

		// Save to database
		if err := r.DB.Create(&schedule).Error; err != nil {
			fmt.Printf("Row %d: failed to save schedule: %v\n", i+1, err)
			return err
		}
	}

	return nil
}

func (r *ScheduleRepository) GetScheduleById(scheduleId uint) (*models.Schedule, error) {
	var schedule models.Schedule
	if err := r.DB.Preload("Room").Preload("User").First(&schedule, scheduleId).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *ScheduleRepository) GetcountScheduleUser() ([]*utils.ScheduleUserCount, error) {
	var counts []*utils.ScheduleUserCount
	if err := r.DB.Table("schedule as s").
		Select("u.full_name AS name, COUNT(s.id) AS count").
		Joins("LEFT JOIN \"user\" AS u ON u.id = s.user_id").
		Group("u.full_name").
		Scan(&counts).Error; err != nil {
		return nil, err
	}
	return counts, nil
}

func (r *ScheduleRepository) GetCountScheduleRoom() ([]*utils.ScheduleRoomCount, error) {
	var counts []*utils.ScheduleRoomCount
	if err := r.DB.Table("schedule").
		Select("room.id as room_id, room.name as room_name, COUNT(schedule.id) as schedule_count").
		Joins("JOIN room ON schedule.room_id = room.id").
		Group("room.id, room.name").
		Scan(&counts).Error; err != nil {
		return nil, err
	}
	return counts, nil

}

func (r *ScheduleRepository) CreateSchedule(createScheduleDto *schedule.CreateRoomScheduleDto) (*models.Schedule, error) {
	newSchedule := &models.Schedule{
		UserID:      createScheduleDto.UserID,
		RoomID:      createScheduleDto.RoomID,
		StartTime:   createScheduleDto.StartTime,
		EndTime:     createScheduleDto.EndTime,
		Status:      models.ScheduleStatus(createScheduleDto.Status),
		Description: createScheduleDto.Description,
		Title:       createScheduleDto.Title,
	}

	if err := r.DB.Create(newSchedule).Error; err != nil {
		return nil, err
	}

	return newSchedule, nil
}

func (r *ScheduleRepository) UpdateSchedule(roomScheduleId uint, dto schedule.UpdateRoomSchedule) (*models.Schedule, error) {
	var existingSchedule models.Schedule
	if err := r.DB.Table("schedule").Where("id = ?", roomScheduleId).First(&existingSchedule).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"status":      models.ScheduleStatus(dto.Status),
		"room_id":     dto.RoomID,
		"start_time":  dto.StartTime,
		"end_time":    dto.EndTime,
		"user_id":     dto.UserID,
		"description": dto.Description,
		"title":       dto.Title,
	}

	fmt.Println(dto)

	if err := r.DB.Table("schedule").Where("id = ?", roomScheduleId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := r.DB.First(&existingSchedule, roomScheduleId).Error; err != nil {
		return nil, err
	}
	return &existingSchedule, nil
}

func (r *ScheduleRepository) DeleteSchedule(roomScheduleId uint) error {
	result := r.DB.Table("schedule").Where("id = ?", roomScheduleId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func containsRole(roles []string, role string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}
func (r *ScheduleRepository) GetAllSchedules(roomId uint, userId uint, roles []string) ([]*models.Schedule, error) {
	var roomSchedules []*models.Schedule

	query := r.DB.Table("schedule").Preload("User").Preload("Room")

	if !containsRole(roles, "admin") {
		query = query.Where("user_id = ?", userId)
	}

	if roomId != 0 {
		query = query.Where("room_id = ?", roomId)
	}

	err := query.Find(&roomSchedules).Error
	if err != nil {
		log.Printf("Error retrieving schedules: %v", err)
		return nil, err
	}

	log.Printf("Retrieved %d schedules", len(roomSchedules))
	return roomSchedules, nil
}

func NewScheduleRepository(db *gorm.DB) Repository.RoomScheduleRepositoryInterface {
	return &ScheduleRepository{
		DB: db,
	}
}
