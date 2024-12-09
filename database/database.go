package database

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/d11m08y03/CC-EOY/config"
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
	if executeSQLFile("./database/create_students.sql", "students") {
		loadStudentDataFromCSV(config.StudentCSVPath)
	}

	if executeSQLFile("./database/create_emails.sql", "emails") {
		loadEmailDataFromCSV(config.CCEmailCSVPath)
	}

	executeSQLFile("./database/create_organisors.sql", "organisors")
}

func executeSQLFile(filename string, tableName string) bool {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read SQL file '%s': %v", filename, err)
	}

	if tableExists(tableName) {
		log.Printf("Table '%s' already exists.\n", tableName)
		return false
	}

	_, err = DB.Exec(string(content))
	if err != nil {
		log.Fatalf("Failed to execute SQL file '%s': %v", filename, err)
	}

	log.Printf("Table '%s' was successfully created.\n", tableName)

	return true
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

func loadStudentDataFromCSV(filename string) {
	if !tableExists("students") {
		log.Println("The 'students' table does not exist. Skipping CSV data loading.")
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open CSV file '%s': %v", filename, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file '%s': %v", filename, err)
	}

  recordCount := 1

	// Insert records into the database
	for i, record := range records {
    // Skip the title row
    if i == 0 {
      continue
    }

		if len(record) < 9 {
			log.Printf("Skipping invalid record: %v", record)
			continue
		}

		insertQuery := `
        INSERT INTO students (
            Timestamp, Email, FullName, ProgrammeOfStudy, Faculty, StudentID, Level, ContactNumber, InternshipWork
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`

		_, err := DB.Exec(
			insertQuery,
			strings.TrimSpace(record[0]),
			strings.TrimSpace(record[1]),
			strings.TrimSpace(record[2]),
			strings.TrimSpace(record[3]),
			strings.TrimSpace(record[4]),
			strings.TrimSpace(record[5]),
			strings.TrimSpace(record[6]),
			strings.TrimSpace(record[7]),
			strings.TrimSpace(record[8]),
		)

		if err != nil {
			log.Printf("Failed to insert record '%v': %v", record, err)
		} else {
      recordCount += 1
		}
	}

  log.Printf("Inserted %d students", recordCount)
}

func loadEmailDataFromCSV(filename string) {
	if !tableExists("emails") {
		log.Println("The 'emails' table does not exist. Skipping CSV data loading.")
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open CSV file '%s': %v", filename, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file '%s': %v", filename, err)
	}

	// Insert records into the database
	for i, record := range records {
    // Skip the title row
    if i == 0 {
      continue
    }

		if len(record) < 3 {
			log.Printf("Skipping invalid record: %v", record)
			continue
		}

		insertQuery := `
        INSERT INTO emails (Email, Password, AppPassword) VALUES (?, ?, ?);
    `

		_, err := DB.Exec(
			insertQuery,
			strings.TrimSpace(record[0]),
			strings.TrimSpace(record[1]),
			strings.TrimSpace(record[2]),
		)

		if err != nil {
			log.Printf("Failed to insert record '%v': %v", record, err)
		} else {
			log.Printf("Record inserted successfully: %v", record)
		}
	}
}
