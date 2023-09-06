package handler

type UserResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Team     string `json:"team"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UserRegisterResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}
