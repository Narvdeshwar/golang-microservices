package main

import (
	"log"
	"order-services/db"
	handlers "order-services/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to order DB", err)
	}
	defer database.Close()

	h := &handlers.Handler{DB: database}
	r := gin.Default()

	r.POST("/orders", h.CreateOrder)

	r.GET("/orders", h.GetAllOrder)

	r.Run(":8082")

}
