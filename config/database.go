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
				ID: 2,
			},
		}, ImageUrl: "", Bio: "Admin", Github: "https://github.com/TuqL3", Facebook: "https://www.facebook.com/TuqL3", Instagram: "https://instagram.com/tuq.l3"},
		{Username: "test2", Password: string(hashedPassword), FullName: "Test2", Email: "test2@test.com", Phone: "0386626021", Roles: []models.Role{
			{
				ID: 2,
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
			Name: "Air condition",
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

	fmt.Println("Seeding data completed")

}
