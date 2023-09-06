package data

import (
	"Immersive_dash/features/mentee"
	"errors"

	"gorm.io/gorm"
)

type menteeQuery struct {
	db *gorm.DB
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
