package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB() *sql.DB {
	// Use an in-memory database for testing
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	DB = db // Assign to the global DB variable for testing
	return db
}

func teardownTestDB(db *sql.DB) {
	db.Close()
}

func TestInitDB(t *testing.T) {
	// Setup in-memory DB
	db := setupTestDB()
	defer teardownTestDB(db)

	// Run migrations
	InitDB()

	tables := []string{"users", "students"}
	for _, table := range tables {
		if !tableExists(table) {
			t.Errorf("Table '%s' should exist but does not", table)
		}
	}
}

func TestCreateTables(t *testing.T) {
	// Setup in-memory DB
	db := setupTestDB()
	defer teardownTestDB(db)

	// Run migrations
	CreateTables()

	tables := []string{"users", "students"}
	for _, table := range tables {
		if !tableExists(table) {
			t.Errorf("Table '%s' should exist but does not", table)
		}
	}
}

func TestTableExists(t *testing.T) {
	// Setup in-memory DB
	db := setupTestDB()
	defer teardownTestDB(db)

	// Create a table manually
	_, err := DB.Exec(`CREATE TABLE test_table (id INTEGER PRIMARY KEY)`)
	if err != nil {
		t.Fatalf("Failed to create test table: %v", err)
	}

	if !tableExists("test_table") {
		t.Error("Expected table 'test_table' to exist, but it does not")
	}

	if tableExists("non_existent_table") {
		t.Error("Expected table 'non_existent_table' to not exist, but it does")
	}
}

func TestCreateTable(t *testing.T) {
	// Setup in-memory DB
	db := setupTestDB()
	defer teardownTestDB(db)

	// Run createTable
	createTable("new_table", "CREATE TABLE new_table (id INTEGER PRIMARY KEY)")

	// Verify table exists
	if !tableExists("new_table") {
		t.Error("Expected table 'new_table' to exist, but it does not")
	}
}
