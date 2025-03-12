package database

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

// Connect initializes the database connection
func Connect(connStr string) {
    var err error
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    if err = db.Ping(); err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    fmt.Println("Database connection established")
}

// GetDB returns the database connection
func GetDB() *sql.DB {
    return db
}