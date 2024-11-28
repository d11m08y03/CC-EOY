package main

import (
	"github.com/d11m08y03/CC-EOY/database"
	"github.com/d11m08y03/CC-EOY/routes"
)

func main() {
	database.InitDB()
	router := routes.SetupRouter()
	router.Run(":8080")
}
