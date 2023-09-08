package class

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

type Core struct {
	ID            uint
	Name          string
	Start_Date    time.Time
	Graduate_Date time.Time
}

func ModelToCore(model Class) Core {
	core := Core{
		ID:            model.ID,
		Name:          model.Name,
		Start_Date:    model.Start_Date,
		Graduate_Date: model.Graduate_Date,
	}
	return core
}

type ClassDataInterface interface {
	GetClassByID(id uint) (*Core, error)
	CreateClass(name string, startDate, graduateDate time.Time) (*Core, error)
	UpdateClass(id uint, name string, startDate, graduateDate time.Time) (*Core, error)
	DeleteClassByID(id uint) error
}

type ClassServiceInterface interface {
	GetClassByID(id uint) (*Core, error)
	CreateClass(name string, startDate, graduateDate time.Time) (*Core, error)
	UpdateClass(id uint, name string, startDate, graduateDate time.Time) (*Core, error)
	DeleteClassByID(id uint) error
}

type ClassData struct {
	DB *gorm.DB
}

func NewClassData(db *gorm.DB) ClassDataInterface {
	return &ClassData{
		DB: db,
	}
}

func (c *ClassData) GetClassByID(id uint) (*Core, error) {
	var class Core
	if err := c.DB.First(&class, id).Error; err != nil {
		return nil, err
	}
	return &class, nil
}

func (c *ClassData) CreateClass(name string, startDate, graduateDate time.Time) (*Core, error) {
	class := &Core{
		Name:          name,
		Start_Date:    startDate,
		Graduate_Date: graduateDate,
	}

	if err := c.DB.Create(class).Error; err != nil {
		return nil, err
	}

	return class, nil
}

func (c *ClassData) UpdateClass(id uint, name string, startDate, graduateDate time.Time) (*Core, error) {
	var class Core
	if err := c.DB.First(&class, id).Error; err != nil {
		return nil, err
	}

	class.Name = name
	class.Start_Date = startDate
	class.Graduate_Date = graduateDate

	if err := c.DB.Save(&class).Error; err != nil {
		return nil, err
	}

	return &class, nil
}

func (c *ClassData) DeleteClassByID(id uint) error {
	return c.DB.Delete(&Core{}, id).Error
}

type ClassService struct {
	Data ClassDataInterface
}

func NewClassService(data ClassDataInterface) ClassServiceInterface {
	return &ClassService{
		Data: data,
	}
}

func (cs *ClassService) GetClassByID(id uint) (*Core, error) {
	return cs.Data.GetClassByID(id)
}

func (cs *ClassService) CreateClass(name string, startDate, graduateDate time.Time) (*Core, error) {
	return cs.Data.CreateClass(name, startDate, graduateDate)
}

func (cs *ClassService) UpdateClass(id uint, name string, startDate, graduateDate time.Time) (*Core, error) {
	return cs.Data.UpdateClass(id, name, startDate, graduateDate)
}

func (cs *ClassService) DeleteClassByID(id uint) error {
	return cs.Data.DeleteClassByID(id)
}
