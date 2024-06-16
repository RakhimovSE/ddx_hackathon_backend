package routes

import (
	"ddx_hackathon_backend/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func setupExerciseRoutes(router *gin.Engine, db *gorm.DB) {
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
}
