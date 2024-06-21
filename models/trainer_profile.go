package models

import "github.com/jinzhu/gorm"

type TrainerProfile struct {
    gorm.Model
    UserID      uint       `json:"userId"`
    Experience  int        `json:"experience"`
    Bio         string     `json:"bio"`
    Specialties []Specialty `gorm:"many2many:trainer_specialties;" json:"specialties"`
}

type TrainerSpecialty struct {
    TrainerProfileID uint `gorm:"primary_key"`
    SpecialtyID      uint `gorm:"primary_key"`
}
