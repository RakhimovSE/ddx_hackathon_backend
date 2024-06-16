package routes

import (
	"ddx_hackathon_backend/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func setupTrainingPlanRoutes(router *gin.Engine, db *gorm.DB) {
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
}
