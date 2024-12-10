package controllers

import (
	"fmt"
	"net/http"

	"github.com/d11m08y03/CC-EOY/auth"
	"github.com/d11m08y03/CC-EOY/config"
	"github.com/d11m08y03/CC-EOY/logger"
	"github.com/d11m08y03/CC-EOY/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	logger.Info("Register controller hit")

	var input models.Organisor
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	input.Password = string(hashedPassword)

	if err := models.CreateOrganisor(input); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Organisor registered successfully"})
}

func Login(c *gin.Context) {
	logger.Info("Login controller hit")

	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.FindUserByEmail(input.Email)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	isAdmin := user.IsAdmin == 1

	token, err := auth.GenerateJWT(user.ID, user.Email, isAdmin)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	logger.Info(fmt.Sprintf("%s logged in successfully", user.Email))

	c.JSON(http.StatusOK, gin.H{"token": token, "is_admin": isAdmin})
}

func CreateAdmin(c *gin.Context) {
	logger.Info("Create Admin controller hit")

	if config.Environment == "Prod" {
		logger.Error("Cannot create admin in development environment")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create admin in development environment"})
		return
	}

	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser, err := models.FindUserByEmail(input.Email)
	if err == nil && existingUser.ID > 0 {
		logger.Error("Email already exists")
		c.JSON(http.StatusConflict, gin.H{"error": "Email is already registered"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Failed to hash password")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	admin := models.Organisor{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		IsAdmin:  1,
	}

	if err := models.CreateOrganisor(admin); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Admin user created successfully"})
}
