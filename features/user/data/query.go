package data

import (
	"Immersive_dash/features/user"
	"errors"

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
	// panic("unimplemented")
	var data User
	tx := repo.db.Where("email = ? and password = ?", email, password).Find(&data)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user.Core{}, errors.New("data not found")
	}
	dataLogin = UserModelToCore(data)
	return dataLogin, nil
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
func (repo *userQuery) Register(input user.Core) (user.Core, error) {
	// panic("unimplemented")
	userGorm := UserCoreToModel(input)
	tx := repo.db.Create(&userGorm) // proses query insert
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	return UserModelToCore(userGorm), nil
}

// Update implements user.UserDataInterface.
func (repo *userQuery) Update(input user.Core) (user.Core, error) {
	// panic("unimplemented")
	userGorm := UserCoreToModel(input)
	tx := repo.db.Model(&User{}).Where("id = ?", userGorm.ID).Updates(userGorm)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	return UserModelToCore(userGorm), nil
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
