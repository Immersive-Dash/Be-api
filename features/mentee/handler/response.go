package handler

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
