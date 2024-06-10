package routes

import (
	"ddx_hackathon_backend/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.Static("/static", "./static")

	router.GET("/users", func(c *gin.Context) {
		handlers.GetUsers(c, db)
	})
	router.POST("/users", func(c *gin.Context) {
		handlers.CreateUser(c, db)
	})
	router.DELETE("/users/:id", func(c *gin.Context) {
		handlers.DeleteUser(c, db)
	})
	router.PATCH("/users/:id", func(c *gin.Context) {
		handlers.UpdateUser(c, db)
	})
	router.GET("/training_plans", func(c *gin.Context) {
		handlers.GetTrainingPlans(c, db)
	})
	router.POST("/training_plans", func(c *gin.Context) {
		handlers.CreateTrainingPlan(c, db)
	})
	router.DELETE("/training_plans/:id", func(c *gin.Context) {
		handlers.DeleteTrainingPlan(c, db)
	})
	router.PATCH("/training_plans/:id", func(c *gin.Context) {
		handlers.UpdateTrainingPlan(c, db)
	})
	router.GET("/workouts", func(c *gin.Context) {
		handlers.GetWorkouts(c, db)
	})
	router.POST("/workouts", func(c *gin.Context) {
		handlers.CreateWorkout(c, db)
	})
	router.POST("/login", func(c *gin.Context) {
		handlers.LoginUser(c, db)
	})
	router.GET("/exercises", func(c *gin.Context) {
		handlers.GetExercises(c, db)
	})
}
