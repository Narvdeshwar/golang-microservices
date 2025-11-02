package main

import (
	"user-services/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/user", handlers.CreateUser)
	r.GET("/users", handlers.GetAllUser)

	r.Run(":8081") // run on port 8081
}
