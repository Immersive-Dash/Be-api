package database

import (
	feedback "Immersive_dash/features/feedback/data"
	mentee "Immersive_dash/features/mentee/data"
	user "Immersive_dash/features/user/data"

	"gorm.io/gorm"
)

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&user.Team{})
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&mentee.Mentee{})
	db.AutoMigrate(&feedback.Feedback{})
	//db.AutoMigrate(&mentee.StatusEntity{})

}
