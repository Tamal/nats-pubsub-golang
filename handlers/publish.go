package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

// Publish - Push messages to NATS
func Publish(c *gin.Context) {
	nc := c.MustGet("nc").(*nats.Conn)
	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))
}
