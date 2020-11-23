package main

import (
	"nats-streaming/handlers"
	"nats-streaming/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	// Connect NATS middleware
	router.Use(middlewares.ConnectNATS())
	router.Use(middlewares.Subscribe())
	// CORS middleware
	router.Use(middlewares.CORSMiddleware())

	// Group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", handlers.Ping)
		v1.POST("/pub", handlers.Publish)
	}

	router.Run()
}
