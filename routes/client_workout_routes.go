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
	router.GET("/training_plans/:training_plan_id/workouts", func(c *gin.Context) {
		handlers.GetWorkoutsByTrainingPlan(c, db)
	})
	router.GET("/client_workouts/:client_workout_id", func(c *gin.Context) {
		handlers.GetClientWorkoutExercises(c, db)
	})
}
