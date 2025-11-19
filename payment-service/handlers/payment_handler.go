package handlers

import (
	"fmt"
	"net/http"
	"os"
	"payment-services/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) MakePayment(ctx *gin.Context) {
	var pay models.Payment
	if err := ctx.BindJSON(&pay); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if pay.OrderID == 0 {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "Order id is required"})
		return
	}

	// Order url creation
	orderURL := fmt.Sprintf("%s/order/%d", os.Getenv("ORDER_SERVICE_URL"), pay.OrderID)

	// getting the order url to check whether the order has been crated or not
	resp, err := http.Get(orderURL)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error connecting to the order services"})
		return
	}

	defer resp.Body.Close()

	err = h.DB.QueryRow("INSERT into payments(order_id,amount,status) values($1,$2,$3) RETURNING id", pay.OrderID, pay.Amount, "SUCCESS").Scan(&pay.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error recordigng the payment"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Payment successful", "data": pay})
}
