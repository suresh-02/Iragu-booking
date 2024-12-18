package models

import "gorm.io/gorm"

type UserCreds struct {
	gorm.Model
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique"`
}
