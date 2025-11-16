package handlers

import (
	"net/http"
	"order-services/models"

	"github.com/gin-gonic/gin"
)

var orders = []models.Order{}

func (h *Handler) CreateOrder(ctx *gin.Context) {
	var newOrder models.Order
	if err := ctx.BindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// !TODO : user validation api required  
	resp, err := http.Get("http://user-service:8081/users")
	if err != nil || resp.StatusCode != 200 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error connecting to the user services"})
	}

	err = h.DB.QueryRow("Insert into orders (user_id,item,amount) values($1,$2,$3) Returning id", newOrder.UserID, newOrder.Item, newOrder.Amount).Scan(&newOrder.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the Order"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order Created Successfully", "data": newOrder})
}

func (h *Handler) GetAllOrder(ctx *gin.Context) {
	rows, err := h.DB.Query("Select id,user_id,item,amount from orders")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	defer rows.Close()
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.Item, &o.Amount); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning the user data"})
			return
		}
		orders = append(orders, o)
	}
	if len(orders) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No Order found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}
