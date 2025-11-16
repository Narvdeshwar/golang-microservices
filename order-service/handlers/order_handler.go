package handlers

import (
	"net/http"
	"order-services/models"

	"github.com/gin-gonic/gin"
)

var orders = []models.Order{}

func (h *Handlers) CreateOrder(ctx *gin.Context) {
	var newOrder models.Order
	if err := ctx.BindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Validating the user by calling the user services
	resp, err := http.Get("http://user-service:8081/users")
	if err != nil || resp.StatusCode != 200 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error connecting to the user services"})
	}

	err = h.DB.QueryRow("Insert in orders (user_id,item,amount) values($1,$2,$3) Returning id", newOrder.UserID, newOrder.Item, newOrder.Amount).Scan(&newOrder.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the Order"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order Created Successfully", "data": newOrder})
}

func GetAllOrder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, orders)
}
