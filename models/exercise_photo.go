package models

import (
	"github.com/jinzhu/gorm"
)

type ExercisePhoto struct {
	gorm.Model
	ExerciseID uint   `json:"exerciseId"`
	PhotoURL   string `json:"photoUrl"`
}