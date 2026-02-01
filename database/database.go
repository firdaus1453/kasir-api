package database

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
    // Fallback: Switch to port 5432 and force SSL if 6543 fails
    connectionString = strings.Replace(connectionString, ":6543", ":5432", 1)

    if !strings.Contains(connectionString, "sslmode=") {
        if strings.Contains(connectionString, "?") {
            connectionString += "&sslmode=require"
        } else {
            connectionString += "?sslmode=require"
        }
    }

	// Open database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Test connection
	err = db.Ping()
	if err != nil {
        // Detailed error logging
        log.Printf("DB Connection Error: %v", err)
		return nil, err
	}

	// Set connection pool settings (optional tapi recommended)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
	return db, nil
}
