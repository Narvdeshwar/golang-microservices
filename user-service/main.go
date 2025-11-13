package main

import (
	"log"
	"user-services/db"
	handlers "user-services/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Error in connecting to DB", err)
	}
	defer database.Close()

	h := handlers.Handler{DB: database}
	r := gin.Default()
	r.POST("/user", h.CreateUser)
	r.GET("/users", h.GetAllUser)
	r.Run(":8081") // run on port 8081
}
