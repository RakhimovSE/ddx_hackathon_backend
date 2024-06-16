package models

import (
	"github.com/jinzhu/gorm"
)

type TrainingPlan struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Workouts    []Workout `gorm:"foreignkey:TrainingPlanID"`
	CreatedByID *uint     `json:"created_by_id"`
	CreatedBy   *User     `json:"created_by"`
}
