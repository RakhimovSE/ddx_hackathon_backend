package models

import (
	"github.com/jinzhu/gorm"
)

type Workout struct {
	gorm.Model
	TrainingPlanID uint              `json:"training_plan_id"`
	Name           string            `json:"name"`
	Description    string            `json:"description"`
	DaysUntilNext  int               `json:"days_until_next"`
	Exercises      []WorkoutExercise `json:"exercises" gorm:"foreignkey:WorkoutID"`
}
