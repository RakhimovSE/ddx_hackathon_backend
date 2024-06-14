package models

import (
	"github.com/jinzhu/gorm"
)

type Exercise struct {
	gorm.Model
	Name             string          `json:"name"`
	Muscles          []Muscle        `gorm:"many2many:exercise_muscles;"`
	AdditionalMuscles []Muscle       `gorm:"many2many:exercise_additional_muscles;"`
	TypeID           uint            `json:"type_id"`
	Type             ExerciseType    `json:"type"`
	Equipments       []Equipment     `gorm:"many2many:exercise_equipments;"`
	DifficultyID     uint            `json:"difficulty_id"`
	Difficulty       Difficulty      `json:"difficulty"`
	Photos           []ExercisePhoto `json:"photos" gorm:"foreignkey:ExerciseID"`
}
