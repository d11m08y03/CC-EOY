package config

import (
	"fmt"
	"os"

	"github.com/d11m08y03/CC-EOY/logger"
	"github.com/joho/godotenv"
)

var Port string
var Environment string

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to load .env file, exiting. Error: %s", err.Error()))
	}

	Port = os.Getenv("PORT")
	Environment = os.Getenv("ENVIRONMENT")
}
