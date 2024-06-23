package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ClientWorkout struct {
    gorm.Model
    ClientTrainingPlanID uint                    `json:"client_training_plan_id"`
    Name                 string                  `json:"name"`
	Description          string                  `json:"description"`
    DaysUntilNext  int               `json:"days_until_next"`
    Order                int                     `json:"order"`
    Exercises            []ClientWorkoutExercise `gorm:"foreignkey:ClientWorkoutID"`
    WorkoutID            uint                    `json:"workout_id"`
    StartDate            *time.Time              `json:"start_date"`
    EndDate              *time.Time              `json:"end_date"`
    PlannedStartDate     *time.Time              `json:"planned_start_date"`
    PlannedEndDate       *time.Time              `json:"planned_end_date"`
}
