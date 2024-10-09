package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		claims, ok := user.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Failed to get user claims",
			})
			c.Abort()
			return
		}
		if claims["role"] != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Admin access required",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
