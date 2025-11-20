package db

import (
	"database/sql"
	"log"
	"time"
)

func ConnectDB() (*sql.DB, error) {
	dsn := "postgres://notification:password@notification-db:5432/notificationdb?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Retry loop in case DB is not ready
	for i := 0; i < 5; i++ {
		if err := db.Ping(); err == nil {
			log.Println("Connected to DB successfully")
			return db, nil
		}
		log.Printf("Waiting for DB to be ready... (%d/5)", i+1)
		time.Sleep(3 * time.Second)
	}

	return nil, err
}
