package data

import (
	"Immersive_dash/features/feedback"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	Notes    string
	MenteeID uint
	Status   string
}

func ModelToCore(dataModel Feedback) feedback.Core {
	return feedback.Core{
		ID:       dataModel.ID,
		Notes:    dataModel.Notes,
		MenteeID: dataModel.MenteeID,
		Status:   dataModel.status,
	}
}

func CoreToModel(dataCore feedback.Core) Feedback {
	return Feedback{
		Model:    gorm.Model{},
		Notes:    dataCore.Notes,
		MenteeID: dataCore.MenteeID,
		status:   dataCore.Status,
	}
}
