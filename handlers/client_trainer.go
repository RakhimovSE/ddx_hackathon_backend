package handlers

import (
	"ddx_hackathon_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AddClientTrainer(c *gin.Context, db *gorm.DB) {
	var input models.ClientTrainer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client-trainer relationship"})
		return
	}

	c.JSON(http.StatusOK, input)
}

func DeleteClientTrainer(c *gin.Context, db *gorm.DB) {
	var input models.ClientTrainer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("client_id = ? AND trainer_id = ?", input.ClientID, input.TrainerID).Delete(&models.ClientTrainer{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete client-trainer relationship"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Deleted"})
}
