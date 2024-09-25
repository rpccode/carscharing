package models

import "database/sql"

type Vehicle struct {
	ID           int     `json:"id"`
	LicensePlate string  `json:"license_plate"`
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Status       string  `json:"status"` // disponible, reservado, en uso
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

// Actualizar la ubicación del vehículo
func (v *Vehicle) UpdateLocation(db *sql.DB, lat, long float64) error {
	query := `UPDATE vehicles SET latitude = $1, longitude = $2 WHERE id = $3`
	_, err := db.Exec(query, lat, long, v.ID)
	return err
}

// Cambiar el estado del vehículo
func (v *Vehicle) UpdateStatus(db *sql.DB, status string) error {
	query := `UPDATE vehicles SET status = $1 WHERE id = $2`
	_, err := db.Exec(query, status, v.ID)
	return err
}
