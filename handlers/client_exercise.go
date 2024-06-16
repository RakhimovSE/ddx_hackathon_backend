package handlers

import (
	"net/http"

	"ddx_hackathon_backend/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetClientWorkoutExercises(c *gin.Context, db *gorm.DB) {
	clientWorkoutID := c.Param("client_workout_id")

	var clientWorkout models.ClientWorkout
	if err := db.Preload("Exercises.Sets").First(&clientWorkout, clientWorkoutID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch client workout"})
		return
	}

	var clientWorkoutExercises []models.ClientWorkoutExercise
	if err := db.Where("client_workout_id = ?", clientWorkoutID).Order("order").Preload("Sets").Find(&clientWorkoutExercises).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch client workout exercises"})
		return
	}

	c.JSON(http.StatusOK, clientWorkoutExercises)
}
