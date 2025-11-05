package handlers

import (
	"net/http"
	"user-services/models"

	"github.com/gin-gonic/gin"
)

var users = []models.User{}

func CreateUser(ctx *gin.Context) {
	var newUser models.User
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	ctx.JSON(http.StatusOK, newUser)
}

func GetAllUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, users)
}
