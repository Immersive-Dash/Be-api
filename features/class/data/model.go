package models

import (
	"time"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name          string
	Start_Date    time.Time
	Graduate_Date time.Time
}

func (c *Class) Create(db *gorm.DB) error {
	return db.Create(c).Error
}

func GetClassByID(db *gorm.DB, classID uint) (*Class, error) {
	var class Class
	if err := db.First(&class, classID).Error; err != nil {
		return nil, err
	}

	return &class, nil
}
