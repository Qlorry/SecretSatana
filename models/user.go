package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:100;unique;not null"`
	Password     string `gorm:"not null"`
	Participates bool   `gorm:"default:false"`
}
