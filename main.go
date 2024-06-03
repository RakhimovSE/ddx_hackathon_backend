package main

import (
	"log"
	"os"

	"ddx_hackathon_backend/database"
	"ddx_hackathon_backend/routes"
	"ddx_hackathon_backend/scripts"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.InitDatabase()
	defer database.DB.Close()

	// Check if the argument is "load_data"
	if len(os.Args) > 1 && os.Args[1] == "load_data" {
		scripts.LoadDataFromFile(database.DB)
		return
	}

	router := gin.Default()
	routes.SetupRoutes(router, database.DB)

	router.Run(":8080")
}
