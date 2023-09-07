package data

import (
	"Immersive_dash/features/feedback"
	"Immersive_dash/features/mentee"
	_dataMentee "Immersive_dash/features/mentee/data"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	Notes    string
	MenteeID uint
	status   string
	Mentee   _dataMentee.Mentee
}

func ModelToCore(dataModel Feedback) feedback.Core {
	return feedback.Core{
		ID:       dataModel.ID,
		Notes:    dataModel.Notes,
		MenteeID: dataModel.MenteeID,
		Mentee: mentee.Core{
			FullName: dataModel.Mentee.FullName,
		},
		Status: dataModel.status,
	}
}

func CoreToModel(dataCore feedback.Core) Feedback {
	return Feedback{
		Model:    gorm.Model{},
		Notes:    dataCore.Notes,
		MenteeID: dataCore.MenteeID,
		status:   dataCore.Status,
		Mentee: _dataMentee.Mentee{
			Model:    gorm.Model{},
			FullName: dataCore.Mentee.FullName,
		},
	}
}
