package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"user-services/db"
	"user-services/handlers"
)

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Error in connecting to DB")
	}
	defer database.Close()
	r := gin.Default()
	r.POST("/user", handlers.CreateUser)
	r.GET("/users", handlers.GetAllUser)
	r.Run(":8081") // run on port 8081
}
