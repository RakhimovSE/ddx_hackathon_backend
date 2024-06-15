package models

import (
	"github.com/jinzhu/gorm"
)

type Muscle struct {
	gorm.Model
	Name       string       `json:"name"`
}
