package utils

import (
	"github.com/dgrijalva/jwt-go"
	"server/models"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

type Claims struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}
