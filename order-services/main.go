package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

var orders = []Order{}

func main() {
	r := gin.Default()

	r.POST("/order", func(ctx *gin.Context) {
		var newOrder Order
		fmt.Println(ctx.Request)
		if err := ctx.BindJSON(&newOrder); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		newOrder.Id = len(orders) + 1
		orders = append(orders, newOrder)
		ctx.JSON(http.StatusOK, newOrder)
	})

	r.GET("/orders", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, orders)
	})

	r.Run(":8082")

}
