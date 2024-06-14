package models

import (
	"github.com/jinzhu/gorm"
)

type Difficulty struct {
	gorm.Model
	Level     string      `json:"level"`
	Exercises []Exercise  `json:"exercises"`
}
