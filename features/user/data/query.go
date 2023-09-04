package data

import (
	"Immersive_dash/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// DeleteById implements user.UserDataInterface.
func (repo *userQuery) DeleteById(id uint) error {
	panic("unimplemented")
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string, password string) (dataLogin user.Core, err error) {
	panic("unimplemented")
}

// Read implements user.UserDataInterface.
func (repo *userQuery) Read() ([]user.Core, error) {
	// panic("unimplemented")
	var usersData []User
	tx := repo.db.Preload("Team").Find(&usersData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var usersCore []user.Core
	for _, value := range usersData {
		usersCore = append(usersCore, user.Core{
			ID:       value.ID,
			FullName: value.FullName,
			TeamID:   value.TeamID,
			Team: user.TeamCore{
				ID:   value.TeamID,
				Name: value.Team.Name,
			},
			Email: value.Email,
			Role:  value.Role,
		})
	}
	return usersCore, nil
}

// ReadById implements user.UserDataInterface.
func (repo *userQuery) ReadById(id uint) (user.Core, error) {
	panic("unimplemented")
}

// Register implements user.UserDataInterface.
func (repo *userQuery) Register(input user.Core) error {
	panic("unimplemented")
}

// Update implements user.UserDataInterface.
func (repo *userQuery) Update(input user.Core) error {
	panic("unimplemented")
}

// UpdateById implements user.UserDataInterface.
func (repo *userQuery) UpdateById(id uint, input user.Core) error {
	panic("unimplemented")
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}

}
