package db

import (
	"database/sql"
	"log"
	"os"
	"time"
	  _ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	dsn := os.Getenv("AUTH_DB_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println("Error connecting to the db", err)
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
	return db, db.Ping()

}
