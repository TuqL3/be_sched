package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/config"
	"server/controllers"
	"server/repositories"
	"server/routes"
	"server/services"
	"sync"

	"github.com/gorilla/websocket"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


type WSMessage struct {
    ConversationID uint   `json:"conversation_id"`
    SenderID       uint   `json:"sender_id"`
    ReceiverID     uint   `json:"receiver_id"`
    Content        string `json:"content"`
}

type Client struct {
    userID uint
    conn   *websocket.Conn
}

type WSHub struct {
    clients    map[uint]*Client
    mutex      sync.RWMutex
}

func NewWSHub() *WSHub {
    return &WSHub{
        clients: make(map[uint]*Client),
    }
}

func (h *WSHub) Register(userID uint, client *Client) {
    h.mutex.Lock()
    h.clients[userID] = client
    h.mutex.Unlock()
}

func (h *WSHub) Unregister(userID uint) {
    h.mutex.Lock()
    if _, ok := h.clients[userID]; ok {
        delete(h.clients, userID)
    }
    h.mutex.Unlock()
}

func (h *WSHub) SendToUser(userID uint, message []byte) error {
    h.mutex.RLock()
    client, exists := h.clients[userID]
    h.mutex.RUnlock()

    if exists {
        return client.conn.WriteMessage(websocket.TextMessage, message)
    }
    return nil
}

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}


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

	hub := NewWSHub()


    router.GET("/ws/:userId", func(c *gin.Context) {
        userID := c.Param("userId")
        var uID uint
        fmt.Sscanf(userID, "%d", &uID)
        handleWebSocket(c, hub, uID)
    })
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

	
	if err := router.Run(":8081"); err != nil {
		log.Fatal("failed to run app: ", err)
	}
}

func handleWebSocket(c *gin.Context, hub *WSHub, userID uint) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Printf("Failed to upgrade connection: %v", err)
        return
    }

    client := &Client{
        userID: userID,
        conn:   conn,
    }

    hub.Register(userID, client)
    defer hub.Unregister(userID)

    for {
        _, msgBytes, err := conn.ReadMessage()
        if err != nil {
            break
        }

        var msg WSMessage
        if err := json.Unmarshal(msgBytes, &msg); err != nil {
            continue
        }

        hub.SendToUser(msg.ReceiverID, msgBytes)
    }
}
