package service

import (
	"Immersive_dash/features/mentee"
	"errors"

	"github.com/go-playground/validator"
)

type menteeService struct {
	menteeData mentee.MenteeDataInterface
	validate   *validator.Validate
}

// GetAll implements mentee.MenteeServiceInterface.
func (service *menteeService) GetAll() ([]mentee.Core, error) {
	// panic("unimplemented")
	result, err := service.menteeData.Select()
	return result, err
}

func New(repo mentee.MenteeDataInterface) mentee.MenteeServiceInterface {
	return &menteeService{
		menteeData: repo,
		validate:   validator.New(),
	}
}

func (service *menteeService) Create(input mentee.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("validation error" + errValidate.Error())
	}
	err := service.menteeData.Insert(input)
	return err
}

func (service *menteeService) Delete(id uint) error {
	if id == 0 {
		return errors.New("validation error. invalid ID")
	}
	err := service.menteeData.Delete(id)
	return err
}

func (service *menteeService) GetById(id uint) (mentee.Core, error) {
	return service.menteeData.SelectById(id)
}

func (service *menteeService) Update(id uint, newData mentee.Core) error {

	errValidate := service.validate.Struct(newData)
	if errValidate != nil {
		return errors.New("validation error" + errValidate.Error())
	}

	existingData, err := service.menteeData.SelectById(id)
	if err != nil {
		return err
	}

	existingData.FullName = newData.FullName
	existingData.NickName = newData.NickName
	existingData.Email = newData.Email
	existingData.Phone = newData.Phone
	existingData.CurrentAddress = newData.CurrentAddress
	existingData.HomeAddress = newData.HomeAddress
	existingData.Telegram = newData.Telegram
	existingData.Gender = newData.Gender
	existingData.EducationType = newData.EducationType
	existingData.Major = newData.Major
	existingData.Graduate = newData.Graduate
	existingData.Institution = newData.Institution
	existingData.EmergencyName = newData.EmergencyName
	existingData.EmergencyPhone = newData.EmergencyPhone
	existingData.EmergencyStatus = newData.EmergencyStatus

	err = service.menteeData.Update(existingData)

	return err
}
