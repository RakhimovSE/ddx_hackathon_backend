package models

import "github.com/jinzhu/gorm"

type Specialty struct {
    gorm.Model
    Name string `json:"name"`
}
