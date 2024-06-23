package scripts

import (
	"ddx_hackathon_backend/database"
	"fmt"

	"github.com/jinzhu/gorm"
)

func ResetDatabase(db *gorm.DB) {
	db.Exec("DROP TABLE IF EXISTS client_exercise_sets CASCADE")
	db.Exec("DROP TABLE IF EXISTS client_workout_exercises CASCADE")
	db.Exec("DROP TABLE IF EXISTS client_workouts CASCADE")
	db.Exec("DROP TABLE IF EXISTS client_training_plans CASCADE")
	db.Exec("DROP TABLE IF EXISTS client_trainers CASCADE")
	db.Exec("DROP TABLE IF EXISTS reviews CASCADE")
	db.Exec("DROP TABLE IF EXISTS trainer_specialties CASCADE")
	db.Exec("DROP TABLE IF EXISTS trainer_profiles CASCADE")
	db.Exec("DROP TABLE IF EXISTS workout_exercises CASCADE")
	db.Exec("DROP TABLE IF EXISTS workouts CASCADE")
	db.Exec("DROP TABLE IF EXISTS training_plans CASCADE")
	db.Exec("DROP TABLE IF EXISTS exercise_photos CASCADE")
	db.Exec("DROP TABLE IF EXISTS exercise_sets CASCADE")
	db.Exec("DROP TABLE IF EXISTS exercises CASCADE")
	db.Exec("DROP TABLE IF EXISTS equipment CASCADE")
	db.Exec("DROP TABLE IF EXISTS exercise_types CASCADE")
	db.Exec("DROP TABLE IF EXISTS muscles CASCADE")
	db.Exec("DROP TABLE IF EXISTS difficulties CASCADE")
	db.Exec("DROP TABLE IF EXISTS specialties CASCADE")
	db.Exec("DROP TABLE IF EXISTS users CASCADE")

	database.MigrateDB(db)

	fmt.Println("Database reset successfully")
}
