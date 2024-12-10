package repositories

import (
	"fmt"
	"log"
	"server/models"
	"server/utils"
	"time"

	"gorm.io/gorm"
)

func getSchedulesToNotify(db *gorm.DB) ([]models.Schedule, error) {
	var schedules []models.Schedule
	now := time.Now()
	notifyTime := now.Add(2 * time.Hour)

	err := db.Where("start_time BETWEEN ? AND ? AND status = ?", now, notifyTime, models.SchedulePending).
		Find(&schedules).Error
	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func NotifyUsers(db *gorm.DB) {
	schedules, err := getSchedulesToNotify(db)
	if err != nil {
		log.Printf("Error fetching schedules: %v", err)
		return
	}

	for _, schedule := range schedules {
		// Lấy thông tin email từ user liên kết
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

			// Cập nhật trạng thái sau khi thông báo
			schedule.Status = models.ScheduleResolve
			db.Save(&schedule)
		}
	}
}
