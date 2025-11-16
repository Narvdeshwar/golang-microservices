package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func ConnectDB() (*sql.DB, error) {
	connectionStr := "postgres://order:password@order-db:5432/orderdb?sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
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
