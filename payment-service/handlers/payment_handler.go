package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"payment-services/models"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
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
	orderURL := fmt.Sprintf("%s/orders/%d", os.Getenv("ORDER_SERVICE_URL"), pay.OrderID)

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

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Println("Failed to connect the rabbit mq", err)
		return
	}

	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open channel:", err)
		return
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare("Payment notification", true, false, false, false, nil)
	if err != nil {
		log.Println("Failed to declare the Queue", err)
		return
	}

	msg := fmt.Sprintf("Payment successfully recevied %d", pay.OrderID)

	err = ch.Publish("", queue.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(msg)})
	if err != nil {
		log.Println("Error in publishing the channel", err)
		return
	}

	log.Println("Notification sent to queue:", msg)
	ctx.JSON(http.StatusOK, gin.H{"message": "Payment successful", "data": pay})
}
