package database

import (
	"ddx_hackathon_backend/models"

	"github.com/jinzhu/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.TrainingPlan{},
		&models.Workout{},
		&models.WorkoutExercise{},
		&models.ExerciseSet{},
		&models.Exercise{},
		&models.ExercisePhoto{},
		&models.Muscle{},
		&models.Equipment{},
		&models.ExerciseType{},
		&models.Difficulty{},
		&models.TrainerProfile{},
		&models.Review{},
		&models.Specialty{},
	)
}
