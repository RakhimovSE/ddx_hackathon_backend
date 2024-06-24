package handlers

import (
	"net/http"
	"strconv"

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

func GetClientsByTrainerID(c *gin.Context, db *gorm.DB) {
	trainerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trainer ID"})
		return
	}

	var clientTrainerRelations []models.ClientTrainer
	if err := db.Where("trainer_id = ?", trainerID).Find(&clientTrainerRelations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch client-trainer relations"})
		return
	}

	var clientIDs []uint
	for _, relation := range clientTrainerRelations {
		clientIDs = append(clientIDs, relation.ClientID)
	}

	var clients []models.User
	if err := db.Where("id IN (?)", clientIDs).Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}

	c.JSON(http.StatusOK, clients)
}
