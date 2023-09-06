package handler

type MenteeResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	Telegram string `json:"telegram"`
