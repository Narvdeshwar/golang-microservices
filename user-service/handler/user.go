package handlers

import (
	"net/http"
	"user-services/models"

	"github.com/gin-gonic/gin"
)

var users = []models.User{}

func (h *Handler) CreateUser(ctx *gin.Context) {
	var newUser models.User
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.DB.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
		newUser.Name, newUser.Email,
	).Scan(&newUser.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully", "data": newUser})
}

func (h *Handler) GetAllUser(ctx *gin.Context) {
	rows, err := h.DB.Query("Select *from users")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch the user"})
		return
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Email, &u.Name); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning the user data"})
			return
		}
		users = append(users, u)
	}
	if len(users) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No users found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}
