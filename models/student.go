package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/d11m08y03/CC-EOY/database"
	"github.com/d11m08y03/CC-EOY/logger"
)

type Student struct {
	Timestamp        sql.NullString `json:"timestamp"`
	Email            sql.NullString `json:"email"`
	FullName         sql.NullString `json:"full_name"`
	ProgrammeOfStudy sql.NullString `json:"programme_of_study"`
	Faculty          sql.NullString `json:"faculty"`
	StudentID        string         `json:"student_id"`
	Level            sql.NullString `json:"level"`
	ContactNumber    sql.NullString `json:"contact_number"`
	InternshipWork   sql.NullString `json:"internship_work"`
	Presence         bool           `json:"presence"`
	OrganiserID      string         `json:"organiser_id"`
}

type FindPartialStudentDB struct {
	FullName  sql.NullString `json:"full_name"`
	StudentID string         `json:"student_id"`
}

type CreateStudentPayload struct {
	StudentID string `json:"student_id"`
	FullName  string `json:"full_name"`
}

type MarkStudentPresentPayload struct {
	StudentID string `json:"student_id"`
}

func (s *Student) Create(payload CreateStudentPayload, organiserID string) error {
	var existingStudentID int
	err := database.DB.QueryRow("SELECT StudentID FROM students WHERE StudentID = ?", payload.StudentID).Scan(&existingStudentID)
	if err != sql.ErrNoRows {
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to query if student %s is already present: %s", payload.StudentID, err.Error()))
		}

		if err == nil {
			logger.Info(fmt.Sprintf("Attempt to duplicate a student (%s) by organisor %s", payload.StudentID, organiserID))
			return errors.New("Student with this ID already exists")
		}

		return err
	}

	query := `
    INSERT INTO students (StudentID, FullName, OrganiserID)
    VALUES (?, ?, ?);`

	_, err = database.DB.Exec(query, payload.StudentID, payload.FullName, 1, organiserID)
	if err != nil {
		return err
	}

	s.Timestamp.String = time.Now().Format("2006-01-02 15:04:05")
	s.StudentID = payload.StudentID
	s.FullName.String = payload.FullName
	s.OrganiserID = organiserID
	s.Presence = true

	return nil
}

func MarkAsPresent(payload MarkStudentPresentPayload, organiserID string) error {
	var existingStudentID int
	var currentPresence bool
	err := database.DB.QueryRow("SELECT StudentID, Presence FROM students WHERE StudentID = ?", payload.StudentID).Scan(&existingStudentID, &currentPresence)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Student not found")
		}

		return err
	}

	if currentPresence {
		return fmt.Errorf("Student is already marked as present")
	}

	updateQuery := `
		UPDATE students
		SET Presence = ?, OrganiserID = ?
		WHERE StudentID = ?;`

	_, err = database.DB.Exec(updateQuery, true, organiserID, payload.StudentID)
	if err != nil {
		return fmt.Errorf("Failed to mark student as present: %v", err)
	}

	return nil
}

func GetPartialStudentByID(studentID string) (*FindPartialStudentDB, error) {
	query := "SELECT FullName, StudentID FROM students WHERE StudentID = ?"

	var student FindPartialStudentDB
	err := database.DB.QueryRow(query, studentID).Scan(
		&student.FullName,
		&student.StudentID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("student with StudentID %s not found", studentID)
		}
		log.Println(err.Error())
		return nil, fmt.Errorf("error retrieving student: %v", err)
	}

	return &student, nil
}

func GetFullStudentByID(studentID string) (*Student, error) {
	query := "SELECT Timestamp, Email, FullName, ProgrammeOfStudy, Faculty, StudentID, Level, ContactNumber, InternshipWork FROM students WHERE StudentID = ?"

	row := database.DB.QueryRow(query, studentID)

	var student Student

	err := row.Scan(
		&student.Timestamp,
		&student.Email,
		&student.FullName,
		&student.ProgrammeOfStudy,
		&student.Faculty,
		&student.StudentID,
		&student.Level,
		&student.ContactNumber,
		&student.InternshipWork,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Student with StudentID %s not found", studentID)
		}
		log.Println(err.Error())
		return nil, fmt.Errorf("Error retrieving student: %v", err)
	}

	return &student, nil
}
