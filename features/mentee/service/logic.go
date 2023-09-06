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
