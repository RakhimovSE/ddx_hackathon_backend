package models

import (
	"github.com/jinzhu/gorm"
)

type Muscle struct {
	gorm.Model
	Name       string       `json:"name"`
	Exercises  []Exercise   `gorm:"many2many:exercise_muscles;"`
	Exercises2 []Exercise   `gorm:"many2many:exercise_additional_muscles;"`
}
