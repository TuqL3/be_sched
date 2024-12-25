package repositories

import (
	"fmt"
	"log"
	"server/models"
	"server/utils"
	"time"

	"gorm.io/gorm"
)

func GetSchedulesWithinThreeHours(db *gorm.DB) ([]models.Schedule, error) {
	var schedules []models.Schedule
	now := time.Now()
	twoHoursLater := now.Add(2 * time.Hour)
	twoHoursOneMinuteLater := twoHoursLater.Add(1 * time.Minute)

	err := db.Preload("User").
		Where("start_time BETWEEN ? AND ? AND status = ?", twoHoursLater, twoHoursOneMinuteLater, models.ScheduleResolve).
		Find(&schedules).Error
	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func NotifyUsers(db *gorm.DB) {
	schedules, err := GetSchedulesWithinThreeHours(db)
	if err != nil {
		log.Printf("Error fetching schedules: %v", err)
		return
	}

	for _, schedule := range schedules {
		userEmail := schedule.User.Email

		subject := fmt.Sprintf("Thông báo: %s sắp bắt đầu", schedule.Title)
		body := fmt.Sprintf("Xin chào %s,\n\nSự kiện '%s' sẽ bắt đầu lúc %s. Vui lòng chuẩn bị.",
			schedule.User.FullName,
			schedule.Title,
			schedule.StartTime.Format("15:04 - 02/01/2006"),
		)

		err := utils.SendEmail(userEmail, subject, body)
		if err != nil {
			log.Printf("Error sending email to %s: %v", userEmail, err)
		} else {
			log.Printf("Email sent successfully to %s", userEmail)

			schedule.Status = models.ScheduleResolve
			db.Save(&schedule)
		}
	}
}
