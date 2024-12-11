package config

import (
	"fmt"
	"os"

	"github.com/d11m08y03/CC-EOY/logger"
	"github.com/joho/godotenv"
)

var Port string
var Environment string
var StudentCSVPath string
var CCEmailCSVPath string
var EmailRecipient string
var JWTKey string

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		logger.Fatal(fmt.Sprintf("Failed to load .env file, exiting. Error: %s", err.Error()))
	}

	missingVars := []string{}

	// Load environment variables
	config := map[string]*string{
		"PORT":              &Port,
		"ENVIRONMENT":       &Environment,
		"STUDENT_CSV_PATH":  &StudentCSVPath,
		"CC_EMAIL_CSV_PATH": &CCEmailCSVPath,
		"EMAIL_RECIPIENT":   &EmailRecipient,
		"JWT_KEY":           &JWTKey,
	}

	for key, target := range config {
		*target = os.Getenv(key)
		if *target == "" {
			missingVars = append(missingVars, key)
		}
	}

	// Check for missing variables
	if len(missingVars) > 0 {
		logger.Fatal(fmt.Sprintf("Missing required environment variables: %v", missingVars))
	}

	// Log non-sensitive configurations
	logger.Info(fmt.Sprintf("Configuration loaded: Port=%s, Environment=%s, StudentCSVPath=%s, CCEmailCSVPath=%s, EmailRecipient=%s",
		Port, Environment, StudentCSVPath, CCEmailCSVPath, EmailRecipient))
}
