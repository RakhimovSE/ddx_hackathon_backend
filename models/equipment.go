package models

import (
	"github.com/jinzhu/gorm"
)

type Equipment struct {
	gorm.Model
	Name      string      `json:"name"`
	Exercises []Exercise  `gorm:"many2many:exercise_equipments;"`
}
