package handlers

import (
	"net/http"
	"order-services/models"
	"order-services/services"

	"github.com/gin-gonic/gin"
)

var orders = []models.Order{}

func CreateOrder(ctx *gin.Context) {
	var newOrder models.Order
	if err := ctx.BindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if !services.UserExits(newOrder.UserID) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User Doesn't exits"})
		return
	}
	newOrder.ID = len(orders) + 1
	orders = append(orders, newOrder)
	ctx.JSON(http.StatusOK, newOrder)
}

func GetAllOrder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, orders)
}
