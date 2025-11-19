package main

import (
	"log"
	"payment-services/db"
	"payment-services/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to DB", err)
	}

	h := &handlers.Handler{DB: database}
	
	r := gin.Default()
	r.POST("/payment", h.MakePayment)
	r.Run(":8083")
}
