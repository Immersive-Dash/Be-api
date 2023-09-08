package handler

import (
	"Immersive_dash/features/feedback"
)

type FeedbackRequest struct {
	Notes    string `json:"notes"`
	MenteeID uint   `json:"id_mentee"`
	Status   string `json:"status"`
}

func RequestToCore(input FeedbackRequest) feedback.Core {
	return feedback.Core{
		Notes:    input.Notes,
		MenteeID: input.MenteeID,
		Status:   input.Status,
	}
}
