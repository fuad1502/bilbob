package main

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"
)

// SafeDB is a wrapper around sql.DB that provides a mutex to make it safe for concurrent use
type SafeDB struct {
	mu sync.Mutex
	db *sql.DB
}

// ConnectDB connects to a Postgres database and checks if the connection is working and returns a SafeDB struct pointer if successful
func ConnectPGDB(host string, user string, password string, dbname string) (*SafeDB, error) {
	connStr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", host, user, password, dbname)
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
