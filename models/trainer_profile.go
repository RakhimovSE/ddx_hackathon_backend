package models

import "github.com/jinzhu/gorm"

type TrainerProfile struct {
    gorm.Model
    UserID      uint       `json:"userId"`
    Experience  int        `json:"experience"`
    Bio         string     `json:"bio"`
    Specialties []Specialty `gorm:"many2many:trainer_specialties;" json:"specialties"`
    User        User       `json:"user"`
}