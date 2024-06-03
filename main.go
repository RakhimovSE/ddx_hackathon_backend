package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"ddx_hackathon_backend/handlers"
	"ddx_hackathon_backend/models"
)

var db *gorm.DB

func main() {
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

    db, err = gorm.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Не удалось подключиться к базе данных: %v", err)
    }
    defer db.Close()

    db.AutoMigrate(&models.User{}, &models.TrainingPlan{})

    router := gin.Default()
    router.GET("/users", func(c *gin.Context) {
        handlers.GetUsers(c, db)
    })
    router.POST("/users", func(c *gin.Context) {
        handlers.CreateUser(c, db)
    })
    router.DELETE("/users/:id", func(c *gin.Context) {
        handlers.DeleteUser(c, db)
    })
    router.PATCH("/users/:id", func(c *gin.Context) {
        handlers.UpdateUser(c, db)
    })
    router.GET("/training_plans", func(c *gin.Context) {
        handlers.GetTrainingPlans(c, db)
    })
    router.POST("/training_plans", func(c *gin.Context) {
        handlers.CreateTrainingPlan(c, db)
    })
    router.DELETE("/training_plans/:id", func(c *gin.Context) {
        handlers.DeleteTrainingPlan(c, db)
    })
    router.PATCH("/training_plans/:id", func(c *gin.Context) {
        handlers.UpdateTrainingPlan(c, db)
    })
    router.POST("/login", func(c *gin.Context) {
        handlers.LoginUser(c, db)
    })

    router.Run(":8080")
}
