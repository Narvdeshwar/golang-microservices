package main

import (
	"notification-services/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/notify", handlers.SendNotification)
	r.Run(":8084")
}
