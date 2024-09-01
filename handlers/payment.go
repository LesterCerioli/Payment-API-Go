package handlers

import (
	"net/http"
	"payment-gateway/services"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	var paymentRequest services.PaymentRequest
	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := services.CreatePaymentService(paymentRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetPaymentStatus(c *gin.Context) {
	paymentID := c.Param("id")
	response, err := services.GetPaymentStatusService(paymentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
