package handler

type MenteeResponse struct {
	ID             uint   `json:"id"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Telegram       string `json:"telegram"`
	Gender         string `json:"gender"`
	CurrentAddress string `json:"current_address"`
	HomeAddress    string `json:"home_address"`
}
