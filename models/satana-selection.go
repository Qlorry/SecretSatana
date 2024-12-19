package models

import "gorm.io/gorm"

type SatanaSelection struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	UserOneID uint `gorm:"not null"`             // UserOne (Secret Satana) must be unique
	UserTwoID uint `gorm:"not null"`             // UserTwo is the recipient
	UserOne   User `gorm:"foreignKey:UserOneID"` // Association with User model
	UserTwo   User `gorm:"foreignKey:UserTwoID"` // Association with User model
}
