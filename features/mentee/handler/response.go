package api

type MenteeResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Telegram string `json:"telegram"`
	Gender   string `json:"gender"`
}
