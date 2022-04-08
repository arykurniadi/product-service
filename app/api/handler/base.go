package handler

import (
	"github.com/gin-gonic/gin"
)

var statusCode = 200

// SetStatusCode -- Set status code
func SetStatusCode(statCode int) int {
	statusCode := statCode
	return statusCode
}

func RespondJSON(c *gin.Context, data interface{}, request interface{}) {
	c.JSON(statusCode, data)
	return
}

func RespondError(c *gin.Context, message string, request interface{}) {
	statusCode := SetStatusCode(400)
	data := gin.H{"status": statusCode, "message": message}
	c.JSON(statusCode, data)
	return
}
