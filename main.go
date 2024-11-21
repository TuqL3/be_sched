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
	socketio "github.com/googollee/go-socket.io"
)

func main() {
	config.LoadEnv()
	config.PostgresConnection()
	config.InitCloudinary()
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

	scheduleRepository := repositories.NewScheduleRepository(config.DB)
	scheduleService := services.NewSheduleService(scheduleRepository)
	scheduleController := controllers.NewRoomScheduleController(scheduleService)

	equipmentRepository := repositories.NewEquipmentRepository(config.DB)
	equipmentService := services.NewEquipmentService(equipmentRepository)
	equipmentController := controllers.NewEquipmentController(equipmentService)

	equipmentTypeRepository := repositories.NewEquipmentTypeRepository(config.DB)
	equipmentTypeService := services.NewEquipmentTypeService(equipmentTypeRepository)
	equipmentTypeController := controllers.NewEquipmentTypeController(equipmentTypeService)

	roleRepository := repositories.NewRoleRepository(config.DB)
	roleService := services.NewRoleService(roleRepository)
	roleController := controllers.NewRoleController(roleService)

	permissionRepository := repositories.NewPermissionRepository(config.DB)
	permissionService := services.NewPermissionService(permissionRepository)
	permissionController := controllers.NewPermissionController(permissionService)

	conversationRepository := repositories.NewConversationRepository(config.DB)
	conversationService := services.NewConversationService(conversationRepository)
	conversationController := controllers.NewConversationController(conversationService)

	messageRepository := repositories.NewMessageRepository(config.DB)
	messageService := services.NewMessageService(messageRepository)
	messegeController := controllers.NewMessageController(messageService)

	routes.UserRoute(router, userController)
	routes.RoomRoute(router, roomController)
	routes.ReportRoute(router, reportController)
	routes.RoomScheduleRoute(router, scheduleController)
	routes.EquipmentRoute(router, equipmentController)
	routes.EquipmentTypeRoute(router, equipmentTypeController)
	routes.RoleRoute(router, roleController)
	routes.PermissionRoute(router, permissionController)
	routes.ConversationRoute(router, conversationController)
	routes.MessageRoute(router, messegeController)

	server := socketio.NewServer(nil)
	server.OnEvent("/", "send_message", func(s socketio.Conn, msg map[string]interface{}) {
		server.BroadcastToRoom("/", msg["conversation_id"].(string), "receive_message", msg)
	})

	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))

	go server.Serve()
	defer server.Close()

	if err := router.Run(":8081"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
