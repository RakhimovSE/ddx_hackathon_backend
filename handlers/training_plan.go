package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"ddx_hackathon_backend/models"
)

func GetTrainingPlans(c *gin.Context, db *gorm.DB) {
    var plans []models.TrainingPlan
    if err := db.Find(&plans).Error; err != nil {
        c.AbortWithStatus(http.StatusNotFound)
    } else {
        c.JSON(http.StatusOK, plans)
    }
}

func CreateTrainingPlan(c *gin.Context, db *gorm.DB) {
    var plan models.TrainingPlan
    c.BindJSON(&plan)
    db.Create(&plan)
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
        c.BindJSON(&plan)
        db.Save(&plan)
        c.JSON(http.StatusOK, plan)
    }
}