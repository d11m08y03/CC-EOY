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
	createStudentsTable := `
    CREATE TABLE IF NOT EXISTS students (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        student_number TEXT NOT NULL UNIQUE,
        created_by INTEGER NOT NULL,
        FOREIGN KEY (created_by) REFERENCES users (id) ON DELETE CASCADE
    );`

	createUsersTable := `
      CREATE TABLE users (
          id INTEGER PRIMARY KEY AUTOINCREMENT,
          name TEXT NOT NULL,
          email TEXT NOT NULL UNIQUE,
          password TEXT NOT NULL
      );`

	createTable("users", createUsersTable)
	createTable("students", createStudentsTable)
}

func createTable(table string, cmd string) {
	if tableExists(table) {
		log.Printf("Table '%s' already exists.\n", table)
		return
	}

	_, err := DB.Exec(cmd)
	if err != nil {
		log.Fatalf("Failed to create table '%s': %v", table, err)
	}

	log.Printf("Table '%s' was successfully created.\n", table)
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
