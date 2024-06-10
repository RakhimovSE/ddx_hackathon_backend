package models

import (
	"github.com/jinzhu/gorm"
)

type ExerciseSet struct {
	gorm.Model
	WorkoutExerciseID uint   `json:"workout_exercise_id"`
	Reps              int    `json:"reps"`
	Duration          int    `json:"duration"`
	RestTime          int    `json:"rest_time"`
}
