package tests

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"ddx_hackathon_backend/database"
)

func SetupTestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic("failed to connect database")
	}
	database.MigrateDB(db)
	return db
}
