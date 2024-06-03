package main

import (
	"log"

	"ddx_hackathon_backend/database"
	"ddx_hackathon_backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := database.SetupDatabase()
	defer db.Close()

	router := gin.Default()
	routes.SetupRoutes(router, db)

	router.Run(":8080")
}
