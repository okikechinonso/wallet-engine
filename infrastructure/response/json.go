package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, message string, status int, data interface{}) {
	responsedata := gin.H{
		"message":   message,
		"data":      data,
		"status":    http.StatusText(status),
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(status, responsedata)
}
