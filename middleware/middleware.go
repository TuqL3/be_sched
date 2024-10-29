package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func RolePermissionMiddleware(requiredRoles []string, requiredPermissions []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to parse token claims"})
			c.Abort()
			return
		}

		c.Set("user", claims)

		userRoles, _ := claims["roles"].([]interface{})
		userPermissions, _ := claims["permissions"].([]interface{})

		if !checkRolesPermissions(userRoles, userPermissions, requiredRoles, requiredPermissions) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient role or permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func checkRolesPermissions(userRoles []interface{}, userPermissions []interface{}, requiredRoles []string, requiredPermissions []string) bool {
	for _, role := range userRoles {
		if contains(requiredRoles, role.(string)) {
			for _, permission := range userPermissions {
				if contains(requiredPermissions, permission.(string)) {
					return true
				}
			}
		}
	}
	return false
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
