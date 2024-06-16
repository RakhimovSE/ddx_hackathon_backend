package handlers

import (
	"net/http"
	"time"

	"ddx_hackathon_backend/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetClientWorkouts(c *gin.Context, db *gorm.DB) {
	clientID := c.Param("client_id")
	status := c.DefaultQuery("status", "all")
	dateParam := c.Query("date")

	var workouts []models.ClientWorkout
	query := db.Joins("JOIN client_training_plans ON client_training_plans.id = client_workouts.client_training_plan_id").
		Where("client_training_plans.user_id = ?", clientID)

	if dateParam != "" {
		date, err := time.Parse("2006-01-02", dateParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
			return
		}
		query = query.Where("client_workouts.start_date >= ? AND client_workouts.start_date < ?", date, date.AddDate(0, 0, 1))
	}

	switch status {
	case "upcoming":
		query = query.Where("client_workouts.start_date >= ?", time.Now())
	case "past":
		query = query.Where("client_workouts.start_date < ?", time.Now())
	}

	query = query.Order("client_workouts.start_date").Find(&workouts)
	if query.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch workouts"})
		return
	}

	c.JSON(http.StatusOK, workouts)
}
