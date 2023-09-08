package data

import (
	"Immersive_dash/features/class"
	"time"

	"gorm.io/gorm"
)

type ClassQuery struct {
	db *gorm.DB
}

func NewClassQuery(db *gorm.DB) class.ClassDataInterface {
	return &ClassQuery{
		db: db,
	}
}

// CreateClass creates a new class record.
func (cq *ClassQuery) CreateClass(name string, startDate time.Time, graduateDate time.Time) (*class.Core, error) {
	classModel := class.Class{
		Name:          name,
		Start_Date:    startDate,
		Graduate_Date: graduateDate,
	}
	if err := cq.db.Create(&classModel).Error; err != nil {
		return nil, err
	}

	core := class.ModelToCore(classModel)
	return &core, nil
}

// DeleteClassByID.
func (cq *ClassQuery) DeleteClassByID(id uint) error {
	classModel := class.Class{}
	classModel.ID = id
	if err := cq.db.Delete(&classModel).Error; err != nil {
		return err
	}
	return nil
}

// GetClassByID.
func (cq *ClassQuery) GetClassByID(id uint) (*class.Core, error) {
	classModel := class.Class{}
	if err := cq.db.First(&classModel, id).Error; err != nil {
		return nil, err
	}

	core := class.ModelToCore(classModel)
	return &core, nil
}

// UpdateClass.
func (cq *ClassQuery) UpdateClass(id uint, name string, startDate time.Time, graduateDate time.Time) (*class.Core, error) {
	classModel := class.Class{}
	classModel.ID = id

	if err := cq.db.First(&classModel, id).Error; err != nil {
		return nil, err
	}

	classModel.Name = name
	classModel.Start_Date = startDate
	classModel.Graduate_Date = graduateDate

	if err := cq.db.Save(&classModel).Error; err != nil {
		return nil, err
	}

	core := class.ModelToCore(classModel)
	return &core, nil
}
