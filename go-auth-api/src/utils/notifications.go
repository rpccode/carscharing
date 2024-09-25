package utils

import (
	"go-auth-api/src/config"
	"go-auth-api/src/models"
)

// Enviar recordatorio de devolución
func SendReturnReminder(userID int) error {
	message := "Recuerda devolver el vehículo a tiempo para evitar cargos adicionales."
	notification := models.Notification{
		UserID:  userID,
		Message: message,
	}
	return notification.Send(config.DB)
}
