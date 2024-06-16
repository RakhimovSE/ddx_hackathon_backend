package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ClientWorkoutExercise struct {
    gorm.Model
    ClientWorkoutID    uint                  `json:"client_workout_id"`
    WorkoutExerciseID  uint                  `json:"workout_exercise_id"`
    WorkoutExercise    WorkoutExercise       `json:"workout_exercise"`
    Sets               []ClientExerciseSet   `gorm:"foreignkey:ClientWorkoutExerciseID"`
    Order              int                   `json:"order"`
    StartDate          *time.Time            `json:"start_date"`
    EndDate            *time.Time            `json:"end_date"`
    PlannedStartDate   *time.Time            `json:"planned_start_date"`
    PlannedEndDate     *time.Time            `json:"planned_end_date"`
}
