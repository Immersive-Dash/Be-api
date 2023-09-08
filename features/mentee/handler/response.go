package handler

import "Immersive_dash/features/feedback"

type MenteeResponse struct {
	ID            uint   `json:"id"`
	Class         string `json:"class"`
	FullName      string `json:"full_name"`
	NickName      string `json:"nick_name"`
	Status        string `json:"status"`
	EducationType string `json:"education_type"`
	Gender        string `json:"gender"`
	Email         string `json:"email,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Telegram      string `json:"telegram,omitempty"`
}

type MenteeFeedbackResponse struct {
	ID        uint            `json:"id_mentee"`
	Name      string          `json:"name"`
	Feedbacks []feedback.Core `json:"feedbacks"`
}
