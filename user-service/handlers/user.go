package handlers

import (
	"net/http"
	"user-services/db"
	"user-services/models"

	"github.com/gin-gonic/gin"
)

var users = []models.User{}

func CreateUser(ctx *gin.Context) {
	var newUser models.User
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	database, _ := db.ConnectDB()
	defer database.Close()

	_, err := database.Exec("Insert into users (name,email) values($1,$2)", newUser.Name, newUser.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully", "data": newUser})
}

func GetAllUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, users)
}
