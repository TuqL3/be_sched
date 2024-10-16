package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"server/config"
	"server/controllers"
	"server/repositories"
	"server/routes"
	"server/services"
)

func main() {
	config.LoadEnv()
	config.PostgresConnection()
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	userRepository := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	roomRepository := repositories.NewRoomRepository(config.DB)
	roomService := services.NewRoomService(roomRepository)
	roomController := controllers.NewRoomController(roomService)

	equipmentRepository := repositories.NewEquipmentRepository(config.DB)
	equipmentService := services.NewEquipmentService(equipmentRepository)
	equipmentController := controllers.NewEquipmentController(equipmentService)

	reportRepository := repositories.NewReportRepository(config.DB)
	reportService := services.NewReportService(reportRepository)
	reportController := controllers.NewReportController(reportService)

	roomScheduleRepository := repositories.NewRoomScheduleRepository(config.DB)
	roomScheduleService := services.NewRoomSheduleService(roomScheduleRepository)
	roomScheduleController := controllers.NewRoomScheduleController(roomScheduleService)

	computerRepository := repositories.NewComputerRepository(config.DB)
	computerService := services.NewComputerService(computerRepository)
	computerController := controllers.NewComputerController(computerService)

	airConditionRepository := repositories.NewAirConditionRepository(config.DB)
	airConditionService := services.NewAirConditionService(airConditionRepository)
	airConditionController := controllers.NewAirConditionController(airConditionService)

	tAndChRepository := repositories.NewTAndChRepository(config.DB)
	tAndChService := services.NewTAndChService(tAndChRepository)
	tAndChController := controllers.NewTAndChController(tAndChService)

	routes.UserRoute(router, userController)
	routes.RoomRoute(router, roomController)
	routes.EquipmentRoute(router, equipmentController)
	routes.ReportRoute(router, reportController)
	routes.RoomScheduleRoute(router, roomScheduleController)
	routes.ComputerRoute(router, computerController)
	routes.AirConditionRoute(router, airConditionController)
	routes.TAndChtRoute(router, tAndChController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
