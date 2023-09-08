package service

import (
	"Immersive_dash/app/middlewares"
	"Immersive_dash/features/user"

	"github.com/go-playground/validator"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

// CreateUser implements user.UserServiceInterface.
func (service *userService) CreateUser(input user.Core) (user.Core, error) {
	// panic("unimplemented")
	result, err := service.userData.Register(input)
	if err != nil {
		return user.Core{}, err
	}
	return result, err
}

// DeleteUserById implements user.UserServiceInterface.
func (service *userService) DeleteUserById(id uint) error {
	// panic("unimplemented")
	err := service.userData.DeleteById(id)
	return err
}

// GetUser implements user.UserServiceInterface.
func (service *userService) GetUser() ([]user.Core, error) {
	// panic("unimplemented")
	result, err := service.userData.Read()
	if err != nil {
		return []user.Core{}, err
	}
	return result, err
}

// LoginUser implements user.UserServiceInterface.
func (service *userService) LoginUser(email string, password string) (dataLogin user.Core, token string, err error) {
	// panic("unimplemented")
	dataLogin, err = service.userData.Login(email, password)
	if err != nil {
		return user.Core{}, "", err
	}
	token, err = middlewares.CreateToken(dataLogin.Role, dataLogin.ID)
	if err != nil {
		return user.Core{}, "", err
	}
	return dataLogin, token, nil
}

// ReadUserById implements user.UserServiceInterface.
func (service *userService) ReadUserById(id uint) (user.Core, error) {
	// panic("unimplemented")
	result, err := service.userData.ReadById(id)
	if err != nil {
		return user.Core{}, err
	}
	return result, nil
}

// UpdateUser implements user.UserServiceInterface.
func (service *userService) UpdateUser(input user.Core) (user.Core, error) {
	// panic("unimplemented")
	result, err := service.userData.Update(input)
	if err != nil {
		return user.Core{}, err
	}
	return result, nil
}

// UpdateUserById implements user.UserServiceInterface.
func (service *userService) UpdateUserById(id uint, input user.Core) (user.Core, error) {
	// panic("unimplemented")
	result, err := service.userData.UpdateById(id, input)
	if err != nil {
		return user.Core{}, err
	}
	return result, nil
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}
