package models

import (
	"database/sql"
	"errors"

	"github.com/d11m08y03/CC-EOY/database"
)

type Student struct {
	ID            int    `json:"id"`
	StudentNumber string `json:"student_number"`
	CreatedBy     int    `json:"created_by"`
}

func (s *Student) Create(studentNumber string, userID uint) error {
	var existingStudentID int
	err := database.DB.QueryRow("SELECT id FROM students WHERE student_number = ?", studentNumber).Scan(&existingStudentID)
	if err != sql.ErrNoRows {
		if err == nil {
			return errors.New("student with this number already exists")
		}
		return err
	}

	query := `
      INSERT INTO students (student_number, created_by)
      VALUES (?, ?);`

	result, err := database.DB.Exec(query, studentNumber, userID)
	if err != nil {
		return err
	}

	// Get the ID of the newly inserted student
	studentID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	s.ID = int(studentID)
	s.StudentNumber = studentNumber
	s.CreatedBy = int(userID)

	return nil
}
