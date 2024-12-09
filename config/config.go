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

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to load .env file, exiting. Error: %s", err.Error()))
	}

	Port = os.Getenv("PORT")
	Environment = os.Getenv("ENVIRONMENT")
	StudentCSVPath = os.Getenv("STUDENT_CSV_PATH")
	CCEmailCSVPath = os.Getenv("CC_EMAIL_CSV_PATH")
	EmailRecipient = os.Getenv("EMAIL_RECIPIENT")

	logger.Info(fmt.Sprintf("Loaded port: %s", Port))
	logger.Info(fmt.Sprintf("Loaded environment: %s", Environment))
	logger.Info(fmt.Sprintf("Loaded student csv file path: %s", StudentCSVPath))
	logger.Info(fmt.Sprintf("Loaded cc email csv file path: %s", CCEmailCSVPath))
	logger.Info(fmt.Sprintf("Loaded email recipient: %s", EmailRecipient))
}
