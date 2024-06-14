package models

import (
	"github.com/jinzhu/gorm"
)

type ExerciseType struct {
	gorm.Model
	Name      string      `json:"name"`
	Exercises []Exercise  `json:"exercises"`
}
