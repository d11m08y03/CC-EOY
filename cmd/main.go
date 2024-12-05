package main

import (
	"github.com/d11m08y03/CC-EOY/config"
	"github.com/d11m08y03/CC-EOY/database"
	"github.com/d11m08y03/CC-EOY/logger"
	"github.com/d11m08y03/CC-EOY/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	if config.Environment == "Prod" {
		logger.StartFileLogging()
		gin.SetMode(gin.ReleaseMode)
	}

	database.InitDB()
	router := routes.SetupRouter()
	router.Run(config.Port)

	if config.Environment == "Prod" {
		logger.StopFileLogging()
	}
}
