package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/d11m08y03/CC-EOY/database"
)

type Student struct {
	ID               int    `json:"id"`
	Email            string `json:"email"`
	FullName         string `json:"full_name"`
	ProgrammeOfStudy string `json:"programme_of_study"`
	Faculty          string `json:"faculty"`
	StudentID        string `json:"student_id"`
	Level            string `json:"level"`
	ContactNumber    string `json:"contact_number"`
	InternshipWork   string `json:"internship_work"`
	Presence         bool   `json:"presence"`
	OrganiserID      string `json:"organiser_id"`
}

type CreateStudentPayload struct {
	StudentID   string `json:"student_id"`
	FullName    string `json:"full_name"`
	OrganisorID string `json:"organiser_id"`
}

func (s *Student) Create(payload CreateStudentPayload) error {
	var existingStudentID int
	err := database.DB.QueryRow("SELECT StudentID FROM students WHERE StudentID = ?", payload.StudentID).Scan(&existingStudentID)
	if err != sql.ErrNoRows {
		if err == nil {
			return errors.New("student with this ID already exists")
		}
		return err
	}

	query := `
    INSERT INTO students (StudentID, FullName, Presence, OrganiserID)
    VALUES (?, ?, ?, ?);`

	result, err := database.DB.Exec(query, payload.StudentID, payload.FullName, 1, payload.OrganisorID)
	if err != nil {
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	s.ID = int(lastInsertID)
	s.StudentID = payload.StudentID
	s.FullName = payload.FullName
	s.OrganiserID = payload.OrganisorID
	s.Presence = true

	return nil
}

func MarkAsPresent(studentID int, organiserID int) error {
	var existingStudentID int
	err := database.DB.QueryRow("SELECT id FROM students WHERE id = ?", studentID).Scan(&existingStudentID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("student not found")
		}
		return err
	}

	updateQuery := `
		UPDATE students
		SET Present = ?, OrganiserID = ?
		WHERE StudentID = ?;`

	_, err = database.DB.Exec(updateQuery, true, organiserID, studentID)
	if err != nil {
		return fmt.Errorf("failed to mark student as present: %v", err)
	}

	return nil
}
