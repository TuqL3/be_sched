package utils

import (
	"github.com/dgrijalva/jwt-go"
)

type Response struct {
	Status  uint        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type Claims struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type ReportRoomCount struct {
	RoomName    string `json:"room"`
	ReportCount int    `json:"count"`
}

type ScheduleRoomCount struct {
	RoomName      string `json:"room"`
	ScheduleCount int    `json:"count"`
}

type ScheduleUserCount struct {
	UserName      string `json:"name"`
	Schedulecount int    `json:"count"`
}
