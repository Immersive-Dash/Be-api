package data

import (
	"Immersive_dash/features/mentee"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	FullName        string
	NickName        string
	Email           string `gorm:"unique"`
	Phone           string `gorm:"unique"`
	CurrentAddress  string
	HomeAddress     string
	Telegram        string `gorm:"unique"`
	ClassID         uint
	Class           string
	StatusID        uint
	Gender          string
	EducationType   string
	Major           string
	Graduate        string
	Institution     string
	EmergencyName   string `gorm:"unique"`
	EmergencyPhone  string `gorm:"unique"`
	EmergencyStatus string `gorm:"unique"`
	Status          string
}

// type Class struct {
// 	gorm.Model
// 	Name string
// }

// type StatusEntity struct {
// 	gorm.Model
// 	Name        string
// }

func CoreToModel(coreMentee mentee.Core) Mentee {
	modelMentee := Mentee{
		Model:           gorm.Model{},
		FullName:        coreMentee.FullName,
		NickName:        coreMentee.NickName,
		Email:           coreMentee.Email,
		Phone:           coreMentee.Phone,
		CurrentAddress:  coreMentee.CurrentAddress,
		HomeAddress:     coreMentee.HomeAddress,
		Telegram:        coreMentee.Telegram,
		ClassID:         coreMentee.ClassID,
		Class:           coreMentee.Class,
		Gender:          coreMentee.Gender,
		EducationType:   coreMentee.EducationType,
		Major:           coreMentee.Major,
		Graduate:        coreMentee.Graduate,
		Institution:     coreMentee.Institution,
		EmergencyName:   coreMentee.EmergencyName,
		EmergencyPhone:  coreMentee.EmergencyPhone,
		EmergencyStatus: coreMentee.EmergencyStatus,
		Status:          coreMentee.Status,
	}

	return modelMentee
}

func ModelToCore(modelMentee Mentee) mentee.Core {
	coreMentee := mentee.Core{
		ID:              modelMentee.ID,
		FullName:        modelMentee.FullName,
		NickName:        modelMentee.NickName,
		Email:           modelMentee.Email,
		Phone:           modelMentee.Phone,
		CurrentAddress:  modelMentee.CurrentAddress,
		HomeAddress:     modelMentee.HomeAddress,
		Telegram:        modelMentee.Telegram,
		ClassID:         modelMentee.ClassID,
		Class:           modelMentee.Class,
		Gender:          modelMentee.Gender,
		EducationType:   modelMentee.EducationType,
		Major:           modelMentee.Major,
		Graduate:        modelMentee.Graduate,
		Institution:     modelMentee.Institution,
		EmergencyName:   modelMentee.EmergencyName,
		EmergencyPhone:  modelMentee.EmergencyPhone,
		EmergencyStatus: modelMentee.EmergencyStatus,
		Status:          modelMentee.Status,
	}

	return coreMentee
}
