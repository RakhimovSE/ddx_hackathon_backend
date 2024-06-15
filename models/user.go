package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Name            string          `json:"name"`
    Email           string          `json:"email" gorm:"unique"`
    Password        string          `json:"password"`
    AvatarUrl       *string         `json:"avatarUrl"`
    Role            string          `json:"role" gorm:"default:'client'"`
    TrainerProfile  *TrainerProfile `json:"trainerProfile,omitempty"`
    ReviewsReceived []Review        `json:"reviewsReceived,omitempty" gorm:"foreignkey:TrainerID"`
    ReviewsGiven    []Review        `json:"reviewsGiven,omitempty" gorm:"foreignkey:ClientID"`
}
