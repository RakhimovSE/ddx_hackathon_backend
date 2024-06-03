package models

import "github.com/jinzhu/gorm"

type TrainingPlan struct {
    gorm.Model
    Title   string `json:"title"`
    Content string `json:"content"`
    UserID  uint   `json:"user_id"`
}
