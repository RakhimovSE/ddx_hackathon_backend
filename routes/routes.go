package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./static")

	// User routes
	setupUserRoutes(router, db)
	// Client routes
	setupClientRoutes(router, db)
	// Trainer routes
	setupTrainerRoutes(router, db)
	// Training plan routes
	setupTrainingPlanRoutes(router, db)
	// Workout routes
	setupWorkoutRoutes(router, db)
	// Exercise routes
	setupExerciseRoutes(router, db)
	// Client workout routes
	setupClientWorkoutRoutes(router, db)

	return router
}
