package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConnectDB() (*sql.DB, error) {
	connectionStr := "postgres://user:password@user-db:5432/userdb?sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	log.Println("Connected to DB successfully")
	return db, nil
}
