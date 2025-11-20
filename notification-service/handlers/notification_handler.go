package handlers

import (
	"log"
	"net/http"
	"notification-services/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SendNotification(ctx *gin.Context) {
	var notification models.Notification
	if err := ctx.BindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Now we can safely access fields
	log.Println("📩 Notification Received:")
	log.Println("User ID:", notification.UserId)
	log.Println("Order ID:", notification.OrderId)
	log.Println("Message:", notification.Message)
	log.Println("--------------------------")

	ctx.JSON(http.StatusOK, gin.H{"message": "Notification received Sucessfully."})

}
