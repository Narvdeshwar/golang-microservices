package main

import (
	"notification-services/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	h := &handlers.Handler{}
	r := gin.Default()
	r.POST("/notify", h.SendNotification)
	r.Run(":8084")
}
