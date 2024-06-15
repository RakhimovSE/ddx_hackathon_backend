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

	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "seed_trainers":
			scripts.SeedTrainers(db)
			return
		case "load_data":
			scripts.LoadDataFromFile(db)
			return
		}
	}

	router := routes.SetupRouter(db)
	router.Run(":8080")
}
