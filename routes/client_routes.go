package routes

import (
	"ddx_hackathon_backend/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func setupClientRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/clients/:client_id/trainers", func(c *gin.Context) {
		handlers.GetTrainersForClient(c, db)
	})
	router.POST("/clients/:client_id/trainers", func(c *gin.Context) {
		handlers.AddClientTrainer(c, db)
	})
	router.DELETE("/clients/:client_id/trainers/:trainer_id", func(c *gin.Context) {
		handlers.DeleteClientTrainer(c, db)
	})
	router.GET("/clients/:client_id/training_plans", func(c *gin.Context) {
		handlers.GetClientTrainingPlans(c, db)
	})
}
