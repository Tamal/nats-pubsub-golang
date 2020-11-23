package middlewares

import (
	"nats-streaming/worker"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

// Subscribe - Subscribe to NATS
func Subscribe() gin.HandlerFunc {
	return func(c *gin.Context) {
		nc := c.MustGet("nc").(*nats.Conn)
		nc.Subscribe("foo", func(m *nats.Msg) {
			worker.Worker(string(m.Data))
		})
	}
}
