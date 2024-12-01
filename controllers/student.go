package controllers

import (
	"log"
	"net/http"

	"github.com/d11m08y03/CC-EOY/models"
	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var input struct {
		StudentNumber string `json:"student_number" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := models.Student{}

	err := student.Create(input.StudentNumber, userID.(uint))
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
