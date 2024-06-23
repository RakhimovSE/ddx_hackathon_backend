package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"ddx_hackathon_backend/models"
)

func AddClientTrainer(c *gin.Context, db *gorm.DB) {
	clientID := c.Param("client_id")
	clientIDUint, err := strconv.ParseUint(clientID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}

	var input struct {
		TrainerID uint `json:"trainer_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientTrainer := models.ClientTrainer{
		ClientID:  uint(clientIDUint),
		TrainerID: input.TrainerID,
	}

	if err := db.Create(&clientTrainer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client-trainer relationship"})
		return
	}

	c.JSON(http.StatusOK, clientTrainer)
}

func DeleteClientTrainer(c *gin.Context, db *gorm.DB) {
	clientID := c.Param("client_id")
	trainerID := c.Param("trainer_id")

	clientIDUint, err := strconv.ParseUint(clientID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}

	trainerIDUint, err := strconv.ParseUint(trainerID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trainer ID"})
		return
	}

	if err := db.Where("client_id = ? AND trainer_id = ?", uint(clientIDUint), uint(trainerIDUint)).Delete(&models.ClientTrainer{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete client-trainer relationship"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Deleted"})
}

func GetTrainersForClient(c *gin.Context, db *gorm.DB) {
	clientID := c.Param("client_id")

	var trainers []models.User
	if err := db.Preload("TrainerProfile.Specialties").
		Joins("JOIN client_trainers ON client_trainers.trainer_id = users.id").
		Where("client_trainers.client_id = ? AND client_trainers.deleted_at IS NULL", clientID).
		Where("users.role = ?", "trainer").
		Find(&trainers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch trainers"})
		return
	}

	c.JSON(http.StatusOK, trainers)
}

func GetClientTrainingPlans(c *gin.Context, db *gorm.DB) {
	clientID := c.Param("client_id")
	
	var clientTrainingPlans []models.ClientTrainingPlan
	if err := db.Where("user_id = ? AND deleted_at IS NULL", clientID).Find(&clientTrainingPlans).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch client training plans"})
			return
	}

	c.JSON(http.StatusOK, clientTrainingPlans)
}

func GetClientCompletedExercises(c *gin.Context, db *gorm.DB) {
	clientID := c.Param("client_id")
	exerciseID := c.Param("exercise_id")

	var clientWorkoutExercises []models.ClientWorkoutExercise

	if err := db.Preload("Sets").
		Preload("Exercise").
		Preload("Exercise.Photos").
		Preload("Exercise.Muscles").
		Preload("Exercise.AdditionalMuscles").
		Preload("Exercise.Equipments").
		Preload("Exercise.Type").
		Preload("Exercise.Difficulty").
		Joins("JOIN client_workouts ON client_workouts.id = client_workout_exercises.client_workout_id").
		Joins("JOIN client_training_plans ON client_training_plans.id = client_workouts.client_training_plan_id").
		Where("client_training_plans.user_id = ? AND client_workout_exercises.exercise_id = ?", clientID, exerciseID).
		Order("client_workout_exercises.start_date DESC").
		Find(&clientWorkoutExercises).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch client workout exercises"})
		return
	}

	c.JSON(http.StatusOK, clientWorkoutExercises)
}
