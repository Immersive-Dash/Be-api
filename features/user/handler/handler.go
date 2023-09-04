package handler

import (
	"Immersive_dash/app/helpers"
	"Immersive_dash/features/user"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *userHandler {
	return &userHandler{
		userService: service,
	}
}

func (handler *userHandler) ReadUser(c echo.Context) error {
	result, err := handler.userService.GetUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	var userResponse []UserResponse
	for _, value := range result {
		userResponse = append(userResponse, UserResponse{
			ID:       value.ID,
			FullName: value.FullName,
			Team:     value.Team.Name,
			Email:    value.Email,
		})
		fmt.Println("data: ", userResponse)
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", userResponse))
}
