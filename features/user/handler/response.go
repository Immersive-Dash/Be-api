package handler

type UserResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Team     string `json:"team"`
	Email    string `json:"email"`
}
