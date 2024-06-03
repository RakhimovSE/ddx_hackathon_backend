package models

import "github.com/jinzhu/gorm"

type Exercise struct {
	gorm.Model
	Name            string       `json:"name"`
	Muscle          Muscle       `json:"muscle"`
	AdditionalMuscle Muscle      `json:"additionalMuscle"`
	Type            ExerciseType `json:"type"`
	Equipment       Equipment    `json:"equipment"`
	Difficulty      Difficulty   `json:"difficulty"`
	Photos          []ExercisePhoto `json:"photos" gorm:"foreignKey:ExerciseID"`
}

type ExercisePhoto struct {
	gorm.Model
	ExerciseID uint   `json:"exerciseId"`
	PhotoURL   string `json:"photoUrl"`
}
