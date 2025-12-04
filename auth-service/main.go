package main

import (
	"auth-service/db"
	"auth-service/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		log.Println("Error connecting to database")
		return
	}
	defer database.Close()

	h := &handlers.Handler{DB: database}
	r := gin.Default()

	r.POST("/register", h.RegisterUser)
	r.POST("/login", h.LoginUser)

	r.Run(":8085")

}
