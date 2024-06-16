package models

import (
	"github.com/jinzhu/gorm"
)

type WorkoutExercise struct {
	gorm.Model
	WorkoutID      uint        `json:"workout_id"`
	ExerciseID     uint        `json:"exercise_id"`
	RestTime       int         `json:"rest_time"`
	Order      		 int         `json:"order"`
	Sets           []ExerciseSet `json:"sets" gorm:"foreignkey:WorkoutExerciseID"`
	Exercise       Exercise    `json:"exercise" gorm:"foreignkey:ExerciseID"`
}
