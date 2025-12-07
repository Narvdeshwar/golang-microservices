package main

import (
	// "gateway-service/middleware"
	"gateway-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.Use(middleware.AuthMiddleware())

	r.Any("/api/auth/*any", routes.ReverseProxy("http://auth-service:8085"))
	r.Any("/api/user/*any", routes.ReverseProxy("http://user-service:8081"))
	r.Any("/api/order/*any", routes.ReverseProxy("http://order-service:8082"))
	r.Any("/api/payment/*any", routes.ReverseProxy("https://payment-service:8083"))
	r.Run(":8080")
}
