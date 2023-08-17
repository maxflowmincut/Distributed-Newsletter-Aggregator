package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	
	"newsletter-aggregator/src/db/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Read in the schema from schema.sql
	schemaPath := filepath.Join("..", "sqlite", "schema.sql")
	schemaBytes, err := ioutil.ReadFile(schemaPath)
	if err != nil {
		log.Fatalf("Failed to read schema file: %v", err)
	}
	schema := string(schemaBytes)

	// Connect to the SQLite database
	dbPath := filepath.Join("..", "..", "..", "database", "newsletter-aggregator.db")
	db, err := sqlite.ConnectDB(dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer sqlite.CloseDB(db)

	// Initialize the database using the schema
	if err := sqlite.InitializeDB(db, schema); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	log.Println("Database initialized successfully!")
}