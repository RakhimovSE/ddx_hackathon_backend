package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ClientTrainingPlan struct {
	gorm.Model
	Name               string          `json:"name"`
	Description        string          `json:"description"`
	Workouts           []ClientWorkout `gorm:"foreignkey:ClientTrainingPlanID"`
	CreatedByID        uint            `json:"created_by_id"`
	CreatedBy          User            `json:"created_by" gorm:"foreignkey:CreatedByID"`
	UserID             uint            `json:"user_id"`
	TrainingPlanID     uint            `json:"training_plan_id"`
	StartDate          *time.Time      `json:"start_date"`
	EndDate            *time.Time      `json:"end_date"`
	PlannedStartDate   *time.Time      `json:"planned_start_date"`
	PlannedEndDate     *time.Time      `json:"planned_end_date"`
}
