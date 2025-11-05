package main

import (
	"payment-services/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/payment", handlers.CreatePayment)
	r.GET("/payment", handlers.GetPayment)
	r.Run(":8083")
}
