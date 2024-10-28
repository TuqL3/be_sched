package main

import (
	"log"
	"server/config"
	"server/controllers"
	"server/repositories"
	"server/routes"
	"server/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.PostgresConnection()
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
	}))

	userRepository := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	roomRepository := repositories.NewRoomRepository(config.DB)
	roomService := services.NewRoomService(roomRepository)
	roomController := controllers.NewRoomController(roomService)

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

	categoryRepository := repositories.NewCategoryRepository(config.DB)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryService)

	routes.UserRoute(router, userController)
	routes.RoomRoute(router, roomController)
	routes.ReportRoute(router, reportController)
	routes.RoomScheduleRoute(router, roomScheduleController)
	routes.ComputerRoute(router, computerController)
	routes.AirConditionRoute(router, airConditionController)
	routes.TAndChtRoute(router, tAndChController)
	routes.CategoryRoute(router, categoryController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
