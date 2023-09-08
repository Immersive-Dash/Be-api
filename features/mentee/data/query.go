package data

import (
	_feedback "Immersive_dash/features/feedback"
	"Immersive_dash/features/mentee"
	"errors"

	"gorm.io/gorm"
)

type menteeQuery struct {
	db *gorm.DB
}

// SelectMenteeFeedback implements mentee.MenteeDataInterface.
func (repo *menteeQuery) SelectMenteeFeedback(id uint) ([]mentee.Core, error) {
	// panic("unimplemented")
	var menteeData []Mentee
	var tx *gorm.DB
	//query select
	tx = repo.db.Where("id = ?", id).Preload("Feedbacks").Find(&menteeData)
	if tx.Error != nil {
		return []mentee.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return []mentee.Core{}, errors.New("data not found")
	}

	var menteesCore []mentee.Core
	for _, value := range menteeData {
		var feedbacks []_feedback.Core
		for _, feedback := range value.Feedbacks {
			feedbacks = append(feedbacks, _feedback.Core{
				ID:       feedback.ID,
				Notes:    feedback.Notes,
				MenteeID: feedback.MenteeID,
				Status:   feedback.Status,
			})
		}

		var menteeCore = mentee.Core{
			ID:        value.ID,
			FullName:  value.FullName,
			Feedbacks: feedbacks,
		}
		menteesCore = append(menteesCore, menteeCore)
	}

	return menteesCore, nil
}

// Select implements mentee.MenteeDataInterface.
func (repo *menteeQuery) Select() ([]mentee.Core, error) {
	// panic("unimplemented")
	var menteesData []Mentee
	tx := repo.db.Find(&menteesData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var menteesCore []mentee.Core
	for _, value := range menteesData {
		menteesCore = append(menteesCore, mentee.Core{
			ID:              value.ID,
			FullName:        value.FullName,
			NickName:        value.NickName,
			Email:           value.Email,
			Phone:           value.Phone,
			CurrentAddress:  value.CurrentAddress,
			HomeAddress:     value.HomeAddress,
			Telegram:        value.Telegram,
			ClassID:         value.ClassID,
			Class:           value.Class,
			Gender:          value.Gender,
			EducationType:   value.EducationType,
			Major:           value.Major,
			Graduate:        value.Graduate,
			Institution:     value.Institution,
			EmergencyName:   value.EmergencyName,
			EmergencyPhone:  value.EmergencyPhone,
			EmergencyStatus: value.EmergencyStatus,
			Status:          value.Status,
		})
	}
	return menteesCore, nil
}

func New(db *gorm.DB) mentee.MenteeDataInterface {
	return &menteeQuery{
		db: db,
	}
}

func (repo *menteeQuery) Insert(input mentee.Core) error {

	menteeGorm := CoreToModel(input)
	tx := repo.db.Create(&menteeGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *menteeQuery) Delete(id uint) error {
	var menteeGorm Mentee
	tx := repo.db.Delete(&menteeGorm, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *menteeQuery) SelectById(id uint) (mentee.Core, error) {
	var result Mentee
	tx := repo.db.First(&result, id)
	if tx.Error != nil {
		return mentee.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return mentee.Core{}, errors.New("data not found")
	}

	resultCore := ModelToCore(result)
	return resultCore, nil
}

func (repo *menteeQuery) Update(data mentee.Core) error {
	var existingMentee Mentee
	tx := repo.db.First(&existingMentee, data.ID)
	if tx.Error != nil {
		return tx.Error
	}

	existingMentee.FullName = data.FullName
	existingMentee.NickName = data.NickName
	existingMentee.Email = data.Email
	existingMentee.Phone = data.Phone
	existingMentee.CurrentAddress = data.CurrentAddress
	existingMentee.HomeAddress = data.HomeAddress
	existingMentee.Telegram = data.Telegram
	existingMentee.Gender = data.Gender
	existingMentee.EducationType = data.EducationType
	existingMentee.Major = data.Major
	existingMentee.Graduate = data.Graduate
	existingMentee.Institution = data.Institution
	existingMentee.EmergencyName = data.EmergencyName
	existingMentee.EmergencyPhone = data.EmergencyPhone
	existingMentee.EmergencyStatus = data.EmergencyStatus

	tx = repo.db.Save(&existingMentee)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
