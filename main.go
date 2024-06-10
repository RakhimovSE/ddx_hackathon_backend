package main

import (
	"log"
	"os"

	"ddx_hackathon_backend/database"
	"ddx_hackathon_backend/routes"
	"ddx_hackathon_backend/scripts"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := database.SetupDB()
	defer db.Close()

	// Check if the argument is "load_data"
	if len(os.Args) > 1 && os.Args[1] == "load_data" {
		scripts.LoadDataFromFile(db)
		return
	}

	router := routes.SetupRouter(db)
	router.Run(":8080")
}
