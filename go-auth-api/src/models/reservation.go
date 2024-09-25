package models

import (
	"database/sql"
	"errors"
	"time"
)

type Reservation struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	VehicleID int       `json:"vehicle_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status"` // activa, completada, cancelada
}

// Crear una nueva reserva
func (r *Reservation) Create(db *sql.DB) error {
	// Verificar disponibilidad del vehículo
	query := `SELECT COUNT(*) FROM reservations WHERE vehicle_id = $1 AND status = 'activa' AND 
              ((start_time <= $2 AND end_time >= $2) OR (start_time <= $3 AND end_time >= $3))`
	var count int
	err := db.QueryRow(query, r.VehicleID, r.StartTime, r.EndTime).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("el vehículo no está disponible en el rango de tiempo solicitado")
	}

	// Crear la reserva
	query = `INSERT INTO reservations (user_id, vehicle_id, start_time, end_time, status) 
             VALUES ($1, $2, $3, $4, 'activa') RETURNING id`
	return db.QueryRow(query, r.UserID, r.VehicleID, r.StartTime, r.EndTime).Scan(&r.ID)
}
