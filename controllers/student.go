package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/d11m08y03/CC-EOY/models"
	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	organisorID, exists := c.Get("organisor_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	organisorIDStr := fmt.Sprintf("%v", organisorID)

	var payload models.CreateStudentPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := models.Student{}

	err := student.Create(payload, organisorIDStr)
	if err != nil {
		log.Printf("Error inserting student: %v", err)
		if err.Error() == "student with this number already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "Student with this number already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create student"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"student": student})
}

func MarkStudentAsPresent(c *gin.Context) {
	organisorID, exists := c.Get("organisor_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	organisorIDStr := fmt.Sprintf("%v", organisorID)

	var payload models.MarkStudentPresentPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if payload.StudentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "student_id is required"})
		return
	}

	err := models.MarkAsPresent(payload, organisorIDStr)
	if err != nil {
		log.Printf("Error marking student as present: %v", err)

		if err.Error() == "student not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else if err.Error() == "student is already marked as present" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "student is already marked as present"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to mark student as present"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student marked as present"})
}
