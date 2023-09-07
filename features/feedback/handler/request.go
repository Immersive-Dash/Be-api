package handler

import (
	"Immersive_dash/features/feedback"
	"Immersive_dash/features/mentee"
	_dataMentee "Immersive_dash/features/mentee/data"
)

type FeedbackRequest struct {
	Notes    string             `json:"notes"`
	MenteeID uint               `json:"id_mentee"`
	Mentee   _dataMentee.Mentee `json:"mentee"`
	Status   string             `json:"status"`
}

func RequestToCore(input FeedbackRequest) feedback.Core {
	return feedback.Core{
		Notes:    input.Notes,
		MenteeID: input.MenteeID,
		Mentee: mentee.Core{
			FullName: input.Mentee.FullName,
		},
		Status: input.Status,
	}
}
