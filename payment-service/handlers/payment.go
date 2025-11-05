package handlers

import (
	"net/http"
	"payment-services/models"
	"payment-services/services"

	"github.com/gin-gonic/gin"
)

var payments []models.Payment
var nextPaymentID = 1

func CreatePayment(ctx *gin.Context) {
	var p models.Payment
	if err := ctx.BindJSON(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if !services.OrderExits(p.OrderID) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Order does not exist"})
		return
	}

	p.ID = nextPaymentID
	p.Status = "Success"
	nextPaymentID++
	payments = append(payments, p)

	ctx.JSON(http.StatusOK, p)
}

func GetPayment(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, payments)
}
