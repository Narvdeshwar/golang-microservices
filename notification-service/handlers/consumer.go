package handlers

import (
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func StartRabbitConsumer() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal("RabbitMQ connect error:", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Channel open error:", err)
	}

	q, _ := ch.QueueDeclare("Payment notification", true, false, false, false, nil)
	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	log.Println("Notification Service listening for RabbitMQ messages...")

	for msg := range msgs {
		log.Println("Hi, Ashrith Notification Received:", string(msg.Body))
	}
}
