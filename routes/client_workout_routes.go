package routes

import (
	"ddx_hackathon_backend/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func setupClientWorkoutRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/clients/:client_id/workouts", func(c *gin.Context) {
		handlers.GetClientWorkouts(c, db)
	})
}
