package main

import (
	"log"
	"payment-services/db"
	"payment-services/handlers"
	"payment-services/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to DB", err)
	}

	h := &handlers.Handler{DB: database}

	r := gin.Default()
	r.Use(middleware.AuthMiddleware())
	r.POST("/payment", h.MakePayment)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.Run(":8083")
}
