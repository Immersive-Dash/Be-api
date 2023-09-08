package service

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name         string
	StartDate    time.Time
	GraduateDate time.Time
}

func main() {
	// Open a database connection
	db, err := gorm.Open(sqlite.Open("db_immersive_dash.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate the schema (create the 'classes' table)
	db.AutoMigrate(&Class{})

	// Create a new instance of ClassDAO
	classDAO := NewClassDAO(db)

	// Example usage:
	newClass, err := classDAO.CreateClass("Math 101", time.Now(), time.Now().AddDate(0, 3, 0))
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Class ID:", newClass.ID)
	}
}

type ClassDAO struct {
	DB *gorm.DB
}

func NewClassDAO(db *gorm.DB) *ClassDAO {
	return &ClassDAO{DB: db}
}

func (dao *ClassDAO) CreateClass(name string, startDate, graduateDate time.Time) (*Class, error) {
	class := &Class{
		Name:         name,
		StartDate:    startDate,
		GraduateDate: graduateDate,
	}
	if err := dao.DB.Create(class).Error; err != nil {
		return nil, err
	}
	return class, nil
}

func (dao *ClassDAO) GetClassByID(id uint) (*Class, error) {
	var class Class
	if err := dao.DB.First(&class, id).Error; err != nil {
		return nil, err
	}
	return &class, nil
}

func (dao *ClassDAO) UpdateClass(id uint, name string, startDate, graduateDate time.Time) (*Class, error) {
	var class Class
	if err := dao.DB.First(&class, id).Error; err != nil {
		return nil, err
	}

	class.Name = name
	class.StartDate = startDate
	class.GraduateDate = graduateDate

	if err := dao.DB.Save(&class).Error; err != nil {
		return nil, err
	}
	return &class, nil
}

func (dao *ClassDAO) DeleteClassByID(id uint) error {
	if err := dao.DB.Delete(&Class{}, id).Error; err != nil {
		return err
	}
	return nil
}
