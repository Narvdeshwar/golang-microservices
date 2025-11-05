package main

import (
	"order-services/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/orders", handlers.CreateOrder)

	r.GET("/orders", handlers.GetAllOrder)

	r.Run(":8082")

}
