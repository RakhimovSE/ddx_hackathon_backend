package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"ddx_hackathon_backend/models"
)

func GetWorkouts(c *gin.Context, db *gorm.DB) {
	var workouts []models.Workout
	if err := db.Preload("Exercises").Find(&workouts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch workouts"})
		return
	}
	c.JSON(http.StatusOK, workouts)
}

func CreateWorkout(c *gin.Context, db *gorm.DB) {
	var workout models.Workout
	if err := c.ShouldBindJSON(&workout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&workout).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create workout"})
		return
	}
	c.JSON(http.StatusOK, workout)
}
