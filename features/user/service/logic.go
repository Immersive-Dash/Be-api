package service

import (
	"Immersive_dash/features/user"

	"github.com/go-playground/validator"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

// CreateUser implements user.UserServiceInterface.
func (service *userService) CreateUser(input user.Core) error {
	panic("unimplemented")
}

// DeleteUserById implements user.UserServiceInterface.
func (service *userService) DeleteUserById(id uint) error {
	panic("unimplemented")
}

// GetUser implements user.UserServiceInterface.
func (service *userService) GetUser() ([]user.Core, error) {
	// panic("unimplemented")
	result, err := service.userData.Read()
	return result, err
}

// LoginUser implements user.UserServiceInterface.
func (service *userService) LoginUser(email string, password string) (dataLogin user.Core, token string, err error) {
	panic("unimplemented")
}

// ReadUserById implements user.UserServiceInterface.
func (service *userService) ReadUserById(id uint) (user.Core, error) {
	panic("unimplemented")
}

// UpdateUser implements user.UserServiceInterface.
func (service *userService) UpdateUser(input user.Core) error {
	panic("unimplemented")
}

// UpdateUserById implements user.UserServiceInterface.
func (service *userService) UpdateUserById(id uint, input user.Core) error {
	panic("unimplemented")
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}
