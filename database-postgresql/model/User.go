package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique_index" json:"username"`
	Email string `gorm:"not null;unique_index" json:"email"`
	Company  string `gorm:"not null" json:"company"`
	Profile  string `gorm:"not null" json:"profile"`
	Active bool `gorm:"default:true" json:"active"`
}