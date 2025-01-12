package config

import (
	"fmt"
	"os"
	"server/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	if err := db.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.Report{},
		&models.EquipmentType{},
		&models.Schedule{},
		&models.Role{},
		&models.Permission{},
		&models.Conversation{},
		&models.Message{},
		&models.Equipment{}); err != nil {
		panic(err)
	}

	seedData(db)
}

func seedData(db *gorm.DB) {

	tables := []interface{}{
		&models.Permission{},
		&models.Role{},
		&models.User{},
		&models.EquipmentType{},
		&models.Room{},
		&models.Equipment{},
		&models.Report{},
		&models.Schedule{},
	}

	for _, table := range tables {
		var count int64
		if err := db.Model(table).Count(&count).Error; err != nil {
			panic(fmt.Sprintf("Error checking table: %v", err))
		}

		if count > 0 {
			fmt.Printf("Table %T already has data. Skipping seeding.\n", table)
			return
		}
	}

	fmt.Println("Starting data seeding...")
	permissions := []models.Permission{
		{
			PermissionName: "viewEquipment",
		},
		{
			PermissionName: "createEquipment",
		},
		{
			PermissionName: "modifyEquipment",
		},
		{
			PermissionName: "deleteEquipment",
		},
		{
			PermissionName: "viewEquipmentType",
		},
		{
			PermissionName: "createEquipmentType",
		},
		{
			PermissionName: "modifyEquipmentType",
		},
		{
			PermissionName: "deleteEquipmentType",
		},
		{
			PermissionName: "createPermission",
		},
		{
			PermissionName: "deletePermission",
		},
		{
			PermissionName: "modifyPermission",
		},
		{
			PermissionName: "viewPermission",
		},
		{
			PermissionName: "createReport",
		},
		{
			PermissionName: "modifyReport",
		},
		{
			PermissionName: "deleteReport",
		},
		{
			PermissionName: "viewReport",
		},
		{
			PermissionName: "createRole",
		},
		{
			PermissionName: "deleteRole",
		},
		{
			PermissionName: "modifyRole",
		},
		{
			PermissionName: "viewRole",
		},
		{
			PermissionName: "createRoom",
		},
		{
			PermissionName: "deleteRoom",
		},
		{
			PermissionName: "modifyRoom",
		},
		{
			PermissionName: "viewRoom",
		},
		{
			PermissionName: "createSchedule",
		},
		{
			PermissionName: "modifySchedule",
		},
		{
			PermissionName: "deleteSchedule",
		},
		{
			PermissionName: "viewSchedule",
		},
		{
			PermissionName: "modifyUser",
		},
		{
			PermissionName: "deleteUser",
		},
		{
			PermissionName: "viewUser",
		},
		{
			PermissionName: "viewProfile",
		},
		{
			PermissionName: "createUser",
		},
	}

	for _, permission := range permissions {
		db.Create(&permission)
	}

	roles := []models.Role{
		{RoleName: "admin", Permissions: []models.Permission{
			{ID: 1},
			{ID: 2},
			{ID: 3},
			{ID: 4},
			{ID: 5},
			{ID: 6},
			{ID: 7},
			{ID: 8},
			{ID: 9},
			{ID: 10},
			{ID: 11},
			{ID: 12},
			{ID: 13},
			{ID: 14},
			{ID: 15},
			{ID: 16},
			{ID: 17},
			{ID: 18},
			{ID: 19},
			{ID: 20},
			{ID: 21},
			{ID: 22},
			{ID: 23},
			{ID: 24},
			{ID: 25},
			{ID: 26},
			{ID: 27},
			{ID: 28},
			{ID: 29},
			{ID: 30},
			{ID: 31},
			{ID: 32},
			{ID: 33},
		}},
		{RoleName: "trucban"},
		{RoleName: "giangvien"},
		{RoleName: "giamdoc"},
	}

	for _, role := range roles {
		db.Create(&role)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	if err != nil {
		panic("failed to hash password")
	}
	users := []models.User{
		{Username: "admin", Password: string(hashedPassword), FullName: "Admin", Email: "admin@admin.com", Phone: "0386626021", Roles: []models.Role{
			{
				ID: 1,
			},
		}, ImageUrl: "", Bio: "Admin", Github: "https://github.com/TuqL3", Facebook: "https://www.facebook.com/TuqL3", Instagram: "https://instagram.com/tuq.l3"},
		{Username: "test1", Password: string(hashedPassword), FullName: "Test1", Email: "test1@test.com", Phone: "0386626021", Roles: []models.Role{
			{
				ID: 3,
			},
		}, ImageUrl: "", Bio: "Admin", Github: "https://github.com/TuqL3", Facebook: "https://www.facebook.com/TuqL3", Instagram: "https://instagram.com/tuq.l3"},
		{Username: "test2", Password: string(hashedPassword), FullName: "Test2", Email: "test2@test.com", Phone: "0386626021", Roles: []models.Role{
			{
				ID: 3,
			},
		}, ImageUrl: "", Bio: "Admin", Github: "https://github.com/TuqL3", Facebook: "https://www.facebook.com/TuqL3", Instagram: "https://instagram.com/tuq.l3"},
	}

	for _, user := range users {
		db.Create(&user)
	}

	equipmentTypes := []models.EquipmentType{
		{
			Name: "Computer",
		},
		{
			Name: "Table",
		},
		{
			Name: "Board",
		},
		{
			Name: "Projector",
		},
	}

	for _, equipmenttype := range equipmentTypes {
		db.Create(&equipmenttype)
	}

	rooms := []models.Room{
		{
			Name:   "601",
			Status: "available",
		},
		{
			Name:   "602",
			Status: "occupied",
		},
		{
			Name:   "603",
			Status: "maintenance",
		},
	}

	for _, room := range rooms {
		db.Create(&room)
	}

	equipments := []models.Equipment{
		{
			Name:            "Computer 1",
			EquipmentTypeID: 1,
			RoomID:          1,
			Status:          "working",
		},
		{
			Name:            "Board 1",
			EquipmentTypeID: 4,
			RoomID:          1,
			Status:          "broken",
		},
		{
			Name:            "Table 1",
			EquipmentTypeID: 2,
			RoomID:          1,
			Status:          "maintained",
		},
		{
			Name:            "Computer 1",
			EquipmentTypeID: 1,
			RoomID:          2,
			Status:          "working",
		},
		{
			Name:            "Board 1",
			EquipmentTypeID: 4,
			RoomID:          2,
			Status:          "broken",
		},
		{
			Name:            "Table 1",
			EquipmentTypeID: 2,
			RoomID:          2,
			Status:          "maintained",
		},
		{
			Name:            "Computer 1",
			EquipmentTypeID: 1,
			RoomID:          3,
			Status:          "working",
		},
		{
			Name:            "Board 1",
			EquipmentTypeID: 4,
			RoomID:          3,
			Status:          "broken",
		},
		{
			Name:            "Table 1",
			EquipmentTypeID: 2,
			RoomID:          3,
			Status:          "maintained",
		},
	}

	for _, equipment := range equipments {
		db.Create(&equipment)
	}

	reports := []models.Report{
		{
			RoomID:      1,
			UserID:      2,
			EquipmentID: 1,
			Content:     "Test content",
			Status:      "pending",
		},
		{
			RoomID:      2,
			UserID:      3,
			EquipmentID: 5,
			Content:     "Test content",
			Status:      "pending",
		},
		{
			RoomID:      3,
			UserID:      2,
			EquipmentID: 9,
			Content:     "Test content",
			Status:      "resolve",
		},
		{
			RoomID:      1,
			UserID:      2,
			EquipmentID: 2,
			Content:     "Test content",
			Status:      "pending",
		},
	}

	for _, report := range reports {
		db.Create(&report)
	}

	schedules := []models.Schedule{
		{
			RoomID:      1,
			UserID:      2,
			StartTime:   time.Now().Add(time.Hour * 24),
			EndTime:     time.Now().Add(time.Hour * 26),
			Status:      "pending",
			Description: "Team meeting",
			Title:       "Quarterly Sync",
		},
		{
			RoomID:      2,
			UserID:      3,
			StartTime:   time.Now().Add(time.Hour * 48),
			EndTime:     time.Now().Add(time.Hour * 50),
			Status:      "resolve",
			Description: "Project brainstorming session",
			Title:       "Creative Session",
		},
		{
			RoomID:      3,
			UserID:      3,
			StartTime:   time.Now().Add(time.Hour * 72),
			EndTime:     time.Now().Add(time.Hour * 74),
			Status:      "reject",
			Description: "Client presentation",
			Title:       "Project Alpha Presentation",
		},
	}

	for _, schedule := range schedules {
		db.Create(&schedule)
	}

	fmt.Println("Seeding data completed")

}
