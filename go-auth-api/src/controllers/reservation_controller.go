package controllers

import (
	"go-auth-api/src/config"
	"go-auth-api/src/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Obtener el ID del usuario desde el token JWT
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	reservation.UserID = userID.(int)

	// Establecer tiempo de inicio y fin (puedes cambiar esto según tus necesidades)
	reservation.StartTime = time.Now()
	reservation.EndTime = time.Now().Add(2 * time.Hour)

	if err := reservation.Create(config.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reserva creada exitosamente", "reservation": reservation})
}
