package handlers

import (
	"auth-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) RegisterUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "Invalid user input"})
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	err := h.DB.QueryRow("Insert into users(email,password) values($1,$2) RETURNING id", user.Email, string(hash)).Scan(&user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "User already exits!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user register successfully", "user_id": user.ID})
}
