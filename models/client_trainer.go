package models

import (
	"github.com/jinzhu/gorm"
)

type ClientTrainer struct {
    gorm.Model
    ClientID  uint `json:"client_id"`
    TrainerID uint `json:"trainer_id"`
}
