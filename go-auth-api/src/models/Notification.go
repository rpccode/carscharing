package models

import (
	"database/sql"
	"time"
)

type Notification struct {
	ID      int       `json:"id"`
	UserID  int       `json:"user_id"`
	Message string    `json:"message"`
	SentAt  time.Time `json:"sent_at"`
}

// Enviar notificación
func (n *Notification) Send(db *sql.DB) error {
	query := `INSERT INTO notifications (user_id, message, sent_at) 
              VALUES ($1, $2, $3) RETURNING id`
	return db.QueryRow(query, n.UserID, n.Message, time.Now()).Scan(&n.ID)
}
