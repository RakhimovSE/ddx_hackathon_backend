package handlers

import (
	"net/http"
	"strconv"

	"ddx_hackathon_backend/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetExercises(c *gin.Context, db *gorm.DB) {
	offsetParam := c.DefaultQuery("offset", "0")
	limitParam := c.DefaultQuery("limit", "20")

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
		return
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
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
