package main

import (
	"database/sql"
	"errors"
	"log"
	"sync"
)

// SafeDB is a wrapper around sql.DB that provides a mutex to make it safe for concurrent use
type SafeDB struct {
	mu sync.Mutex
	db *sql.DB
}

func ConnectDB() (*SafeDB, error) {
	log.Println("Connecting to the database...")
	connStr := "host=data user=postgres password=secret dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, errors.New("Failed to connect to the database: " + err.Error())
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, errors.New("Failed to ping the database: " + err.Error())
	}
	return &SafeDB{db: db}, nil
}
