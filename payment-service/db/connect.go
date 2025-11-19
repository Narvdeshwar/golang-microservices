package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func ConnectDB() (*sql.DB, error) {
	dsn := os.Getenv("PAYMENT_DB_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
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
