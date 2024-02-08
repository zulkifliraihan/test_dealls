package models

import (
	"time"

	"gorm.io/gorm"
)

type Swipe struct {
	ID 					uint			`gorm:"primary_key;"`
    SwipingUserID 		uint      		
    SwipedUserID  		uint      		
    Direction      		string   		`gorm:"type:varchar(255); not null"`
	CreatedAt    		time.Time
	UpdatedAt    		time.Time
	DeletedAt 			gorm.DeletedAt	`gorm:"index"`	
	SwipingUser 		User `gorm:"foreignKey:SwipingUserID;references:ID"`
    SwipedUser  		User `gorm:"foreignKey:SwipedUserID;references:ID"`
}