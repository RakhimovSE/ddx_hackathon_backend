package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"ddx_hackathon_backend/models"
)

var DB *gorm.DB

func InitDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbSslMode := os.Getenv("DB_SSLMODE")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		dbHost, dbPort, dbUser, dbName, dbSslMode, dbPassword)

	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB.AutoMigrate(&models.User{}, &models.TrainingPlan{}, &models.Exercise{}, &models.ExercisePhoto{})
}
