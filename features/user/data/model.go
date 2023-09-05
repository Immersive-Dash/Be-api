package data

import (
	"Immersive_dash/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string
	Email    string `gorm:"unique"`
	Password string
	TeamID   uint
	Team     Team
	Role     string
}

func UserModelToCore(dataModel User) user.Core {
	return user.Core{
		ID:       dataModel.ID,
		FullName: dataModel.FullName,
		TeamID:   dataModel.TeamID,
		Team: user.TeamCore{
			ID:   dataModel.TeamID,
			Name: dataModel.Team.Name,
		},
		Email: dataModel.Email,
		Role:  dataModel.Role,
	}
}

func UserCoreToModel(dataCore user.Core) User {
	return User{
		Model:    gorm.Model{},
		FullName: dataCore.FullName,
		Email:    dataCore.Email,
		TeamID:   dataCore.TeamID,
		Team: Team{
			Model: gorm.Model{},
			Name:  dataCore.Team.Name,
		},
		Role: dataCore.Role,
	}
}
