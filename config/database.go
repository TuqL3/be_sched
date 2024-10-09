package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"server/models"
	"time"
)

var DB *gorm.DB

func PostgresConnection() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{NowFunc: func() time.Time {
		return time.Now().UTC()
	}})

	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	start := time.Now()

	for sqlDB.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			fmt.Println("Failed to connection database")
			break
		}
	}

	fmt.Println("Connection to database", sqlDB.Ping() == nil)
	DB = db

	if err := db.AutoMigrate(&models.User{}, &models.Equipment{}, &models.Room{}, &models.Report{}, &models.RoomSchedule{}); err != nil {
		panic(err)
	}
}
