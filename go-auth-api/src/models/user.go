package models

import (
	"database/sql"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"` // Agrega este campo
	PasswordHash string `json:"-"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
}

// Registrar un nuevo usuario
func (u *User) Register(db *sql.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost) // Usa Password aquí
	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, password, email, phone) VALUES ($1, $2, $3, $4) RETURNING id`
	err = db.QueryRow(query, u.Username, string(hashedPassword), u.Email, u.Phone).Scan(&u.ID)
	return err
}

// Autenticar usuario
func (u *User) Authenticate(db *sql.DB, password string) error {
	query := `SELECT id, password_hash FROM users WHERE username = $1`
	err := db.QueryRow(query, u.Username).Scan(&u.ID, &u.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("usuario no encontrado")
		}
		return err
	}
	log.Println("Hash almacenado:", u.PasswordHash)
	log.Println("Contraseña proporcionada:", password)

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	log.Println("Resultado de comparación:", err) // Esto mostrará nil si las contraseñas coinciden
	return err
}
