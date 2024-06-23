package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ClientWorkoutExercise struct {
    gorm.Model
    ClientWorkoutID   uint                  `json:"client_workout_id"`
    ExerciseID        uint                  `json:"exercise_id"`
    RestTime          int                   `json:"rest_time"`
    Order             int                   `json:"order"`
    Sets              []ClientExerciseSet   `gorm:"foreignkey:ClientWorkoutExerciseID"`
    StartDate         *time.Time            `json:"start_date"`
    EndDate           *time.Time            `json:"end_date"`
    PlannedStartDate  *time.Time            `json:"planned_start_date"`
    PlannedEndDate    *time.Time            `json:"planned_end_date"`
    Exercise          Exercise              `json:"exercise" gorm:"foreignkey:ExerciseID"`
}
