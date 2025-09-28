package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	dataDir := "./data"
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create data directory: %v", err)
	}
	db, err := sql.Open("sqlite3", "./data/reviews.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	return db
}
