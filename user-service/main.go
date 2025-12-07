package main

import (
	"log"
	"user-services/db"
	handlers "user-services/handler"
	"user-services/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.ConnectDB()
	log.Print("trying to connect", database, err)
	if err != nil {
		log.Fatal("Error in connecting to DB", err)
		return
	}
	defer database.Close()

	h := handlers.Handler{DB: database}

	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetUserById)
	r.GET("/users", h.GetAllUser)
	r.DELETE("/user/:id", h.DeleteUser)
	r.Run(":8081")
}
