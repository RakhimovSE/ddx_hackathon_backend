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
		&models.ClientTrainer{},
	)

	db.Model(&models.TrainerProfile{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.Review{}).AddForeignKey("client_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.Review{}).AddForeignKey("trainer_id", "users(id)", "CASCADE", "CASCADE")

	db.Model(&models.ClientTrainer{}).AddForeignKey("client_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.ClientTrainer{}).AddForeignKey("trainer_id", "users(id)", "CASCADE", "CASCADE")

	db.Model(&models.Exercise{}).AddForeignKey("created_by_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.Exercise{}).AddForeignKey("type_id", "exercise_types(id)", "RESTRICT", "CASCADE")
	db.Model(&models.Exercise{}).AddForeignKey("difficulty_id", "difficulties(id)", "RESTRICT", "CASCADE")

	db.Model(&models.ExercisePhoto{}).AddForeignKey("exercise_id", "exercises(id)", "CASCADE", "CASCADE")
	db.Model(&models.ExerciseSet{}).AddForeignKey("workout_exercise_id", "workout_exercises(id)", "CASCADE", "CASCADE")
	db.Model(&models.WorkoutExercise{}).AddForeignKey("workout_id", "workouts(id)", "CASCADE", "CASCADE")
	db.Model(&models.WorkoutExercise{}).AddForeignKey("exercise_id", "exercises(id)", "CASCADE", "CASCADE")
	db.Model(&models.Workout{}).AddForeignKey("training_plan_id", "training_plans(id)", "CASCADE", "CASCADE")
	db.Model(&models.TrainingPlan{}).AddForeignKey("created_by_id", "users(id)", "CASCADE", "CASCADE")
}
