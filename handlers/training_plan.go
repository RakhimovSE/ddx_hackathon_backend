package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"ddx_hackathon_backend/models"
)

func GetTrainingPlans(c *gin.Context, db *gorm.DB) {
	var plans []models.TrainingPlan
	if err := db.Preload("Workouts").Find(&plans).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, plans)
	}
}

func CreateTrainingPlan(c *gin.Context, db *gorm.DB) {
	var plan models.TrainingPlan
	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&plan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create training plan"})
		return
	}
	c.JSON(http.StatusOK, plan)
}

func DeleteTrainingPlan(c *gin.Context, db *gorm.DB) {
	id := c.Params.ByName("id")
	var plan models.TrainingPlan
	if err := db.Where("id = ?", id).First(&plan).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		db.Delete(&plan)
		c.JSON(http.StatusOK, gin.H{"id #" + id: "deleted"})
	}
}

func UpdateTrainingPlan(c *gin.Context, db *gorm.DB) {
	id := c.Params.ByName("id")
	var plan models.TrainingPlan
	if err := db.Where("id = ?", id).First(&plan).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		if err := c.ShouldBindJSON(&plan); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Save(&plan)
		c.JSON(http.StatusOK, plan)
	}
}
