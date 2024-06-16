package routes

import (
	"ddx_hackathon_backend/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func setupWorkoutRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/workouts", func(c *gin.Context) {
		handlers.GetWorkouts(c, db)
	})
	router.POST("/workouts", func(c *gin.Context) {
		handlers.CreateWorkout(c, db)
	})
}
