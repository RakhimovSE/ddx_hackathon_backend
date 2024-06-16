package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ClientWorkout struct {
    gorm.Model
    ClientTrainingPlanID uint                    `json:"client_training_plan_id"`
    WorkoutID            uint                    `json:"workout_id"`
    Workout              Workout                 `json:"workout"`
    Exercises            []ClientWorkoutExercise `gorm:"foreignkey:ClientWorkoutID"`
    Order                int                     `json:"order"`
    StartDate            *time.Time              `json:"start_date"`
    EndDate              *time.Time              `json:"end_date"`
    PlannedStartDate     *time.Time              `json:"planned_start_date"`
    PlannedEndDate       *time.Time              `json:"planned_end_date"`
}
