package controllers

import (
	"go-auth-api/src/config"
	"go-auth-api/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProcessReservationPayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := payment.ProcessPayment(config.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar el pago"})
		return
	}

	invoice := models.GenerateInvoice(payment)
	c.JSON(http.StatusOK, gin.H{"message": "Pago procesado", "invoice": invoice})
}
