package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	connStr := "host=localhost port=5432 user=CarSharing dbname=CarSharing sslmode=disable password=5120"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot connect to DB", err)
	}

	fmt.Println("Connected to database!")
}
