package handlers

import (
	"net/http"

	"ddx_hackathon_backend/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetTrainers(c *gin.Context, db *gorm.DB) {
	var trainers []models.User
	if err := db.Preload("TrainerProfile.Specialties").
		Where("role = ?", "trainer").Find(&trainers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch trainers"})
		return
	}

	c.JSON(http.StatusOK, trainers)
}
