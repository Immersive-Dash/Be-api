package data

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string
	Email    string `gorm:"unique"`
	Password string
	TeamID   uint
	Team     Team
	Role     string
}
