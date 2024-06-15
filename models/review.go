package models

import "github.com/jinzhu/gorm"

type Review struct {
    gorm.Model
    ClientID   uint   `json:"clientId"`
    TrainerID  uint   `json:"trainerId"`
    Rating     int    `json:"rating"`
    Comment    string `json:"comment"`
    ReviewType string `json:"reviewType"` // "client_to_trainer" or "trainer_to_client"
    Client     User   `json:"client"`
    Trainer    User   `json:"trainer"`
}
