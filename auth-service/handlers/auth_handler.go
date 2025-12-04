package handlers

import (
	"auth-service/models"
	"auth-service/utils"
	"database/sql"
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

func (h *Handler) LoginUser(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var user models.User

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input data"})
		return
	}
	err := h.DB.QueryRow("Select id,password from users where email=$1", input.Email).Scan(&user.ID, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Email or password"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
		return
	}
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
