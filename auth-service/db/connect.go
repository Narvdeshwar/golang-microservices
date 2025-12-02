package main

import (
	"database/sql"
	"log"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	dsn := os.Getenv("AUTH_DB_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println("Error connecting to the db", err)
		return nil, err
	}
	return db, db.Ping()

}
