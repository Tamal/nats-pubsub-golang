package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

// ConnectNATS - Connect NATS
func ConnectNATS() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Connect NATS
		nc, err := nats.Connect(nats.DefaultURL)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("NATS Connected:", nc.ConnectedUrl())
		c.Set("nc", nc)
		c.Next()
	}
}
