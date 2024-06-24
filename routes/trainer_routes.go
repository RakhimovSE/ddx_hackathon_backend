package routes

import (
	"ddx_hackathon_backend/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func setupTrainerRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/trainers", func(c *gin.Context) {
		handlers.GetTrainers(c, db)
	})

	router.GET("/trainers/:id/clients", func(c *gin.Context) {
		handlers.GetClientsByTrainerID(c, db)
	})
}
