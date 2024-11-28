package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./myapp.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	CreateTables()
}

func CreateTables() {
	tableName := "users"
	if tableExists(tableName) {
		log.Printf("Table '%s' already exists.\n", tableName)
		return
	}

	createUsersTable := `
        CREATE TABLE users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            email TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL
        );`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatalf("Failed to create table '%s': %v", tableName, err)
	}
	log.Printf("Table '%s' was successfully created.\n", tableName)
}

func tableExists(tableName string) bool {
	query := `
    SELECT name
    FROM sqlite_master
    WHERE type = 'table' AND name = ?;`

	var name string

	err := DB.QueryRow(query, tableName).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Fatalf("Error checking if table exists: %v", err)
	}

	return true
}
