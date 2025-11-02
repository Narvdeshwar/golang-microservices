package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{}

func main() {
	r := gin.Default() // middleware ke sath ek router return krta hai

	r.POST("/user", func(ctx *gin.Context) {
		var newUser User
		// fmt.Println("rep", ctx.Data)
		if err := ctx.BindJSON(&newUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newUser.Id = len(users) + 1
		users = append(users, newUser)
		ctx.JSON(http.StatusOK, newUser)
	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, users)
	})
	r.Run(":8080")
}
