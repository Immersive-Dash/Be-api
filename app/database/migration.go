package database

import (
	"Immersive_dash/features/user/data"

	"gorm.io/gorm"
)

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&data.Team{})
	db.AutoMigrate(&data.User{})

}
