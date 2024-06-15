package routes

import (
	"ddx_hackathon_backend/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
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
	router.GET("/clients/:client_id/trainers", func(c *gin.Context) {
		handlers.GetTrainersForClient(c, db)
	})
	router.POST("/clients/:client_id/trainers", func(c *gin.Context) {
		handlers.AddClientTrainer(c, db)
	})
	router.DELETE("/clients/:client_id/trainers/:trainer_id", func(c *gin.Context) {
		handlers.DeleteClientTrainer(c, db)
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
	router.POST("/exercises", func(c *gin.Context) {
		handlers.CreateExercise(c, db)
	})
	router.DELETE("/exercises/:id", func(c *gin.Context) {
		handlers.DeleteExercise(c, db)
	})
	router.PATCH("/exercises/:id", func(c *gin.Context) {
		handlers.UpdateExercise(c, db)
	})

	return router
}
