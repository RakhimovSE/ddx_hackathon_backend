package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ClientTrainingPlan struct {
    gorm.Model
    UserID             uint            `json:"user_id"`
    TrainingPlanID     uint            `json:"training_plan_id"`
    TrainingPlan       TrainingPlan    `json:"training_plan"`
    Name               string          `json:"name"`
	Description        string          `json:"description"`
    Workouts           []ClientWorkout `gorm:"foreignkey:ClientTrainingPlanID"`
    StartDate          *time.Time      `json:"start_date"`
    EndDate            *time.Time      `json:"end_date"`
    PlannedStartDate   *time.Time      `json:"planned_start_date"`
    PlannedEndDate     *time.Time      `json:"planned_end_date"`
}
