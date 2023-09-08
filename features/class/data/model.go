package data

import (
	"Immersive_dash/features/class"
	"time"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name          string
	Start_Date    time.Time
	Graduate_Date time.Time
}

func CoreToModel(coreClass class.Core) Class {
	ModelClass := Class{
		Model:         gorm.Model{},
		Name:          coreClass.Name,
		Start_Date:    coreClass.Start_Date,
		Graduate_Date: coreClass.Graduate_Date,
	}
	return ModelClass
}

func ModelToCore(modelClass Class) class.Core {
	coreClass := class.Core{
		Name:          modelClass.Name,
		Start_Date:    modelClass.Start_Date,
		Graduate_Date: modelClass.Graduate_Date,
	}
	return coreClass
}
