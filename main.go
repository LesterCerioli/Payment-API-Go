package main

import (
	"payment-gateway/handlers"
	"payment-gateway/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middleware.AuthMiddleware())

	payment := router.Group("/payments")
	{
		payment.POST("/create", handlers.CreatePayment)
		payment.GET("/:id", handlers.GetPaymentStatus)
	}

	router.Run(":8080")
}
