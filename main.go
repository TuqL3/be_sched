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

	"github.com/robfig/cron/v3"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WSMessage struct {
	ConversationID uint   `json:"conversation_id"`
	SenderID       uint   `json:"sender_id"`
	ReceiverID     uint   `json:"receiver_id"`
	Content        string `json:"content"`
	CreatedAt      string `json:"created_at"`
}

type Client struct {
	userID uint
	conn   *websocket.Conn
	send   chan []byte
}

type WSHub struct {
	clients    map[uint]*Client
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mutex      sync.RWMutex
}

func NewWSHub() *WSHub {
	return &WSHub{
		clients:    make(map[uint]*Client),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *WSHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mutex.Lock()
			h.clients[client.userID] = client
			h.mutex.Unlock()

		case client := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.clients[client.userID]; ok {
				delete(h.clients, client.userID)
				close(client.send)
			}
			h.mutex.Unlock()

		case message := <-h.broadcast:
			var msg WSMessage
			if err := json.Unmarshal(message, &msg); err != nil {
				log.Printf("Error unmarshaling message: %v", err)
				continue
			}

			h.mutex.RLock()
			if client, ok := h.clients[msg.ReceiverID]; ok {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client.userID)
				}
			}
			h.mutex.RUnlock()
		}
	}
}

func (c *Client) readPump(hub *WSHub) {
	defer func() {
		hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		hub.broadcast <- message
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
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
		send:   make(chan []byte, 256),
	}

	hub.register <- client

	go client.writePump()
	go client.readPump(hub)
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
	go hub.Run()

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

	cron := cron.New()
	cron.AddFunc("@every 1m", func() {
		repositories.NotifyUsers(config.DB)
	})
	cron.Start()

	select {}

	if err := router.Run(":8081"); err != nil {
		log.Fatal("failed to run app: ", err)
	}
}
