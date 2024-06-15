package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"ddx_hackathon_backend/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetExercises(c *gin.Context, db *gorm.DB) {
	offset, limit, err := getPaginationParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var exercises []models.Exercise
	if err := db.Preload("Photos").
		Preload("Muscles").
		Preload("AdditionalMuscles").
		Preload("Equipments").
		Preload("Type").
		Preload("Difficulty").
		Offset(offset).
		Limit(limit).
		Find(&exercises).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch exercises"})
		return
	}
	c.JSON(http.StatusOK, exercises)
}

func CreateExercise(c *gin.Context, db *gorm.DB) {
	var exercise models.Exercise
	if err := c.BindJSON(&exercise); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := db.Create(&exercise).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create exercise"})
		return
	}

	if err := updateExerciseAssociations(db, &exercise); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create exercise associations"})
		return
	}

	c.JSON(http.StatusOK, exercise)
}

func DeleteExercise(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := db.Delete(&models.Exercise{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exercise not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Deleted"})
}

func UpdateExercise(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var exercise models.Exercise
	if err := db.First(&exercise, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exercise not found"})
		return
	}

	var input models.Exercise
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	exercise.Name = input.Name
	exercise.TypeID = input.TypeID
	exercise.DifficultyID = input.DifficultyID

	if err := db.Save(&exercise).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update exercise"})
		return
	}

	if err := updateExerciseAssociationsWithInput(db, &exercise, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update exercise associations"})
		return
	}

	c.JSON(http.StatusOK, exercise)
}

func getPaginationParams(c *gin.Context) (int, int, error) {
	offsetParam := c.DefaultQuery("offset", "0")
	limitParam := c.DefaultQuery("limit", "20")

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid offset parameter")
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid limit parameter")
	}

	return offset, limit, nil
}

func updateExerciseAssociations(db *gorm.DB, exercise *models.Exercise) error {
	if err := db.Model(exercise).Association("Muscles").Replace(exercise.Muscles).Error; err != nil {
		return err
	}
	if err := db.Model(exercise).Association("AdditionalMuscles").Replace(exercise.AdditionalMuscles).Error; err != nil {
		return err
	}
	if err := db.Model(exercise).Association("Equipments").Replace(exercise.Equipments).Error; err != nil {
		return err
	}
	return nil
}

func updateExerciseAssociationsWithInput(db *gorm.DB, exercise, input *models.Exercise) error {
	if err := db.Model(exercise).Association("Muscles").Replace(input.Muscles).Error; err != nil {
		return err
	}
	if err := db.Model(exercise).Association("AdditionalMuscles").Replace(input.AdditionalMuscles).Error; err != nil {
		return err
	}
	if err := db.Model(exercise).Association("Equipments").Replace(input.Equipments).Error; err != nil {
		return err
	}
	return nil
}
