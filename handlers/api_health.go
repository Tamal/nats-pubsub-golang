package handlers

import "github.com/gin-gonic/gin"

// Ping - API Health
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
