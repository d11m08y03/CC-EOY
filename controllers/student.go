package controllers

import (
	"fmt"
	"net/http"

	"github.com/d11m08y03/CC-EOY/email"
	"github.com/d11m08y03/CC-EOY/logger"
	"github.com/d11m08y03/CC-EOY/models"
	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	logger.Info("CreateStudent controller hit")

	organisorID, exists := c.Get("organisor_id")
	if !exists {
		logger.Info("Unauthorized access")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	organisorIDStr := fmt.Sprintf("%v", organisorID)

	var payload models.CreateStudentPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := models.Student{}

	err := student.Create(payload, organisorIDStr)
	if err != nil {
		logger.Error(fmt.Sprintf("Error inserting student %s : %s", payload.FullName, err.Error()))

		if err.Error() == "Student with this ID already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "Student with this ID already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
		}

		return
	}

	logger.Info(fmt.Sprintf("%s marked %s as present", organisorIDStr, payload.StudentID))

	c.JSON(http.StatusOK, gin.H{
		"message":    "Student marked as present",
		"student_id": student.StudentID,
		"full_name":  student.FullName.String,
	})
}

func MarkStudentAsPresent(c *gin.Context) {
	logger.Info("MarkStudentAsPresent controller hit")

	organisorID, exists := c.Get("organisor_id")
	if !exists {
		logger.Info("Unauthorized access")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	organisorIDStr := fmt.Sprintf("%v", organisorID)

	var payload models.MarkStudentPresentPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if payload.StudentID == "" {
		logger.Error("No student ID provided")
		c.JSON(http.StatusBadRequest, gin.H{"error": "student_id is required"})
		return
	}

	err := models.MarkAsPresent(payload, organisorIDStr)
	if err != nil {
		logger.Info(fmt.Sprintf("Error marking student %s as present : %s", payload.StudentID, err.Error()))

		if err.Error() == "Student not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else if err.Error() == "Student is already marked as present" {
			c.JSON(http.StatusConflict, gin.H{"error": "Student is already marked as present"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark student as present"})
		}

		return
	}

	updatedStudent, err := models.GetFullStudentByID(payload.StudentID)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to obtain present student from DB: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to format student details but presence updated"})
		return
	}

	logger.Info(fmt.Sprintf("%s marked %s as present", organisorIDStr, payload.StudentID))
	go email.SendEmail(updatedStudent.Email.String)

	c.JSON(http.StatusOK, gin.H{
		"message":          "Student marked as present",
		"email":            updatedStudent.Email.String,
		"full_name":        updatedStudent.FullName.String,
		"program_of_study": updatedStudent.ProgrammeOfStudy.String,
		"faculty":          updatedStudent.Faculty.String,
		"student_id":       updatedStudent.StudentID,
		"level":            updatedStudent.Level.String,
		"contact_number":   updatedStudent.ContactNumber.String,
		"internship_work":  updatedStudent.InternshipWork.String,
	})
}
