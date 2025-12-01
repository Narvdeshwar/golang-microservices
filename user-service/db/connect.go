package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	connectionStr := "postgres://user:password@user-db:5432/userdb?sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, err
	}

	var pingErr error
	// Retry loop in case DB is not ready
	for i := 0; i < 5; i++ {
		if err := db.Ping(); err == nil {
			log.Println("Connected to DB successfully")
			return db, nil
		}
		log.Printf("Waiting for DB to be ready Attempt (%d/5) %v", i+1, pingErr)
		time.Sleep(3 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect after 5 Attempts in 15 seconds %v", pingErr)
}
