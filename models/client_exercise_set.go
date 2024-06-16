package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ClientExerciseSet struct {
    gorm.Model
    ClientWorkoutExerciseID uint      `json:"client_workout_exercise_id"`
    Reps                    int       `json:"reps"`
    Duration                int       `json:"duration"`
    RestTime                int       `json:"rest_time"`
    Order                   int       `json:"order"`
    StartDate               *time.Time `json:"start_date"`
    EndDate                 *time.Time `json:"end_date"`
    PlannedStartDate        *time.Time `json:"planned_start_date"`
    PlannedEndDate          *time.Time `json:"planned_end_date"`
}
