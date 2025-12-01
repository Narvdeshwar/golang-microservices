package main

import (
	"log"
	"notification-services/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	go handlers.StartRabbitConsumer()

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "notification service running"})
	})
	
	log.Println("Notification service started on port 8084")
	r.Run(":8084")
}
