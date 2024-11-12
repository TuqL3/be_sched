package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FileUploadMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Không tìm thấy file trong request"})
			c.Abort()
			return
		}
		defer file.Close()

		c.Set("file", file)
		c.Set("fileName", header.Filename)

		c.Next()
	}
}
