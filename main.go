package main

import (
	"log"
	"os"

	"ddx_hackathon_backend/database"
	"ddx_hackathon_backend/routes"
	"ddx_hackathon_backend/scripts"

	"github.com/joho/godotenv"
)

func loadEnv() {
  env := os.Getenv("GIN_MODE")
  if env == "release" {
    err := godotenv.Load(".env.release")
    if err != nil {
      log.Fatalf("Error loading .env.release file")
    }
  } else {
    err := godotenv.Load(".env.debug")
    if err != nil {
      log.Fatalf("Error loading .env.debug file")
    }
  }
}

func main() {
	loadEnv()

	db := database.SetupDB()
	defer db.Close()

	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "seed":
			scripts.SeedTrainers(db)
			scripts.SeedTrainingPlans(db)
			scripts.SeedClientTrainers(db)
			scripts.SeedClientTrainingPlans(db)
			return
		case "seed_trainers":
			scripts.SeedTrainers(db)
			return
		case "seed_training_plans":
			scripts.SeedTrainingPlans(db)
			return
		case "seed_client_trainers":
			scripts.SeedClientTrainers(db)
			return
		case "seed_client_training_plans":
			scripts.SeedClientTrainingPlans(db)
			return
		case "load_data":
			scripts.LoadDataFromFile(db)
			return
		}
	}

	router := routes.SetupRouter(db)
	router.Run(":" + os.Getenv("APP_PORT"))
}
