package handler

import "Immersive_dash/features/user"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	TeamID   int    `json:"id_team"`
}

type UserUpdateRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RequestToCore(input UserRequest) user.Core {
	return user.Core{
		FullName: input.FullName,
		TeamID:   uint(input.TeamID),
		Password: input.Password,
		Email:    input.Email,
		Role:     input.Role,
	}
}

func UpdateRequestToCore(input UserUpdateRequest) user.Core {
	return user.Core{
		FullName: input.FullName,
		Password: input.Password,
		Email:    input.Email,
	}

}
