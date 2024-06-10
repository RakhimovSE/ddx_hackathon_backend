package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"ddx_hackathon_backend/handlers"
	"ddx_hackathon_backend/models"
)

func SetupTestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})
	return db
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
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
	return router
}
