package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint        `gorm:"primary_key;"`
	Name            string      `gorm:"type:varchar(255); not null"`
	Email           string      `gorm:"type:varchar(255); unique; not null"`
	Password        string      `gorm:"type:varchar(255); not null"`
	Premium         bool        `gorm:"default:false"`
	LastSwipe       time.Time   `gorm:"default:null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	SwipingUser     []Swipe        `gorm:"foreignKey:SwipingUserID;references:ID"`
	SwipedUser      []Swipe        `gorm:"foreignKey:SwipedUserID;references:ID"`
}
